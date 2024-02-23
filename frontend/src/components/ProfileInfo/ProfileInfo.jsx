import React from "react";
import styles from "./ProfileInfo.module.css";

import TelegramCircle from "../../assets/img/telegramCircle.svg";
import infoCat from "../../assets/img/infoCat.svg";
import infoGift from "../../assets/img/infoGift.svg";

const ProfileInfo = () => {
  return (
    <div className="container">
      <div className={styles.profileInfo_block}>
        <div className={styles.profile_data}>
          <div className={styles.data_left}>
            <div className={styles.data_types}>
              <div className={styles.choose_item}>
                <p>Персональные данные</p>
              </div>
              <div className={styles.choose_item}>
                <p>Карточка питомц</p>
              </div>
            </div>
            <div className={styles.data_main_info}>
              <div className={styles.data_enter}>
                <p>ФИО</p>
                <input type="text" placeholder="ФИО" />
              </div>
              <div className={styles.data_enter}>
                <p>ФИО</p>
                <input type="text" placeholder="ФИО" />
              </div>
              <div className={styles.data_enter}>
                <p>ФИО</p>
                <input type="text" placeholder="ФИО" />
              </div>
              <div className={styles.data_enter}>
                <p>ФИО</p>
                <input type="text" placeholder="ФИО" />
              </div>
              <div className={styles.data_enter}>
                <p>ФИО</p>
                <input type="text" placeholder="ФИО" />
              </div>
              <div className={styles.data_enter}>
                <p>ФИО</p>
                <input type="text" placeholder="ФИО" />
              </div>
            </div>
            <div className={styles.show_contacts}>
                <p>Скрыть контакты</p>
            </div>
            <div className={styles.data_button}>Редактировать</div>
          </div>
          <div className={styles.data_right}>
            <div className={styles.data_photo}>
              <p>Фотография профиля</p>
              <div className={styles.photo_download}>
                <img src="" alt="photo" />
              </div>
            </div>
            <div className={styles.data_can}>
              <p>Мой питомец может сдавать кровь сейчас</p>
              <p>choose</p>
            </div>
            <div className={styles.needBlood}>
              <p>Моему питомцу срочно нужна кровь</p>
              <button>Заполнить форму</button>
            </div>
          </div>
        </div>
        <div className={styles.profile_static}>
          <div className={styles.comeIn}>
            <a href="https://t.me/vetdonor">
              <img src={TelegramCircle} alt="telegram" />
              <p>Присоединяйтесь к нам в Telegram!</p>
            </a>
          </div>
          <div className={styles.interest_info}>
            <p>Интересные статьи</p>
            <div className={styles.interest_items}>
              <div className={styles.interest_item}>
                <a href="#">
                  <img src={infoCat} alt="info" />
                  <div className={styles.interest_item_text}>
                    <p>Как подготовить питомца к сдаче крови?</p>
                  </div>
                </a>
              </div>
              <div className={styles.interest_item}>
                <a href="#">
                  <img src={infoGift} alt="info" />
                  <div className={styles.interest_item_text}>
                    <p>Система бонусов VetDonor</p>
                  </div>
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ProfileInfo;
