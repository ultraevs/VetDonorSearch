import telebot
from telebot import types
import os
import requests
import json
from dotenv import load_dotenv
load_dotenv()
token = os.getenv("TELEGRAM_BOT_TOKEN")

bot = telebot.TeleBot(token='6853137874:AAEQKzX4i9BW5jJRNE2v0toGnv7hq3Wdyuw')

global dt
dt = {
    'id': None,
    'email': None,
    'donationType': None
}

@bot.message_handler(commands=['start'])
def start(message):
    global dt
    types_ = {
        'blood': 'Цельная кровь',
        'plasma': 'Плазма крови',
        'platelets': 'Тромбоциты'
    }

    r = requests.get('https://vetdonor.shmyaks.ru/v1/get_all_donations')

    data = json.loads(r.text)
    if data == None:
        bot.send_message(message.chat.id, 'На данный момент нет активных справок на проверку. Попробуйте позже.')
    else:
        data = data[0]
        photo = requests.get(data['photoPath'])

        c = f'EMAIL: {data["email"]}\nТип: {types_[data["donationType"]]}'

        keyboard = types.InlineKeyboardMarkup()
        buttonOK = types.InlineKeyboardButton(text='Принять✅', callback_data=f'ok')
        buttonCANCEL = types.InlineKeyboardButton(text='Отклонить❌', callback_data=f'cn')
        keyboard.add(buttonOK, buttonCANCEL)

        dt = {'id': data['id'], 'email': data['email'], 'donationType': data['donationType']}

        bot.send_photo(message.chat.id, photo.content, caption=c, reply_markup=keyboard)


@bot.callback_query_handler(func=lambda call: True)
def handler_(call):
    global dt
    if call.data == 'ok':
        bot.send_message(call.from_user.id, 'Справка принята.')
        print(f"id = {dt['id']}")
        print(f"donationType = {dt['donationType']}, email = {dt['email']}")
        r = requests.post('https://vetdonor.shmyaks.ru/v1/delete_application', json={'id': dt['id']})
        print(r)
        print(r.text)
        r = requests.put('https://vetdonor.shmyaks.ru/v1/update_user_stats',
                         json={'donationType': dt['donationType'], 'email': dt['email']})
        print(r)
        print(r.text)
    elif call.data == 'cn':
        bot.send_message(call.from_user.id, 'Справка отклонена.')
        r = requests.post('https://vetdonor.shmyaks.ru/v1/delete_application', json={'id': dt['id']})
        print(r)
        print(r.text)


bot.infinity_polling()