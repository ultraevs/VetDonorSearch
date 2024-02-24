import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.svm import SVC
import pickle
import pymorphy2
import os
import logging
import json
import requests
import sys

local_path = r'/Users/ivanloboda/Documents/code/soba4ki/nn-text/'
server_path = r'internal/api/python-assistant/'

ide_path = server_path

logging.basicConfig(level=logging.CRITICAL)

def train(df):
    X_train, X_test, y_train, y_test = train_test_split(df['text'], df['class'], test_size=0.2, random_state=42)

    tfidf_vectorizer = TfidfVectorizer()
    X_train_tfidf = tfidf_vectorizer.fit_transform(X_train)

    model = SVC(kernel='linear')
    model.fit(X_train_tfidf, y_train)

    # Сохраняем модель и векторизатор с помощью pickle
    with open('model.pkl', 'wb') as model_file:
        pickle.dump(model, model_file)


    with open('vectorizer.pkl', 'wb') as vectorizer_file:
        pickle.dump(tfidf_vectorizer, vectorizer_file)


def extract_city(text):
    df = pd.read_csv(fr'{ide_path}cities.csv')
    cities_list = df.iloc[:, 1].tolist()

    morph_analyzer = pymorphy2.MorphAnalyzer()

    tokens = text.lower().split()

    for token in tokens:
        normal_form = morph_analyzer.parse(token)[0].normal_form
        if normal_form in cities_list:
            return normal_form

    return None

def predict(text, model, vectorizer):
    text_to_predict_tfidf = vectorizer.transform([text])
    predicted_class = model.predict(text_to_predict_tfidf)
    return predicted_class

def load_answers():
    with open(fr'{ide_path}answers.json', 'r', encoding="utf-8") as f:
        return json.load(f)

def main(text, retrain=False):
    """
    A function to process text and provide responses related to blood donation, with an option to retrain the model.
    Parameters:
    - text: the input text to process
    - retrain: a boolean indicating whether to retrain the model (default False)
    Return:
    - ans: a string containing the response related to blood donation based on the input text
    """
    answers = load_answers()
    
    if retrain:
        df = pd.read_csv(f'{ide_path}train_classes.csv')
        train(df)

    # Загрузка модели и векторайзера
    with open(fr'{ide_path}model.pkl', 'rb') as model_file:
        model = pickle.load(model_file)
    with open(fr'{ide_path}vectorizer.pkl', 'rb') as vectorizer_file:
        vectorizer = pickle.load(vectorizer_file)

    class_ = predict(text, model, vectorizer)
    if class_ in ['getblood', 'giveblood']:
        city = extract_city(text)
    if city != None:
        df = pd.read_csv(fr'{ide_path}cities.csv')
        index = df[df.iloc[:, 1] == city].index[0]
        city_code = df.iloc[index, 0]
        is_city = True
    else:
        city_code = None
        is_city = False

    class_ = class_[0]

    all_clinics = requests.get('https://vetdonor.shmyaks.ru/v1/get_all_clinic_cards').json()
    allowed_cards = []

    for clinic in all_clinics:
       if city.lower() in clinic['address'].lower():
           allowed_cards.append(clinic)

    formatted_clinics_text = ''
    for clinic in allowed_cards:
        formatted_clinics_text += f'{clinic["name"]}<br/>{clinic["address"]}<br/>{clinic["contacts"]}<br/>'

    if is_city:
        ans = f'{answers[f"{is_city}"][class_].format(city.capitalize())}<br/><br/>{formatted_clinics_text}'
    else:
        ans = answers[f"{is_city}"][class_]
    return ans


if __name__ == "__main__":
    arg1 = sys.argv[1]
    response = {"callback": main(arg1)}
    
    with open(f"{ide_path}result.json", "w", encoding='utf-8') as file:
        json.dump(response, file, ensure_ascii=False)


    print(1)
    