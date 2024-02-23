import React, { useState } from "react";
import styles from "./ProfileInfo.module.css";

import classNames from "class-names";

import TelegramCircle from "../../assets/img/telegramCircle.svg";
import infoCat from "../../assets/img/infoCat.svg";
import infoGift from "../../assets/img/infoGift.svg";

import RightArrow from "../../assets/img/RightArrow.svg";
import LeftArrow from "../../assets/img/LeftArrow.svg";

const ProfileInfo = () => {
  const [activeItem, setActiveItem] = useState(0);

  const handleItemClick = (item) => {
    setActiveItem(item);
  };
  return (
    <div className="container">
      <div className={styles.profileInfo_block}>
        <div className={styles.profile_data}>
          <div className={styles.data_left}>
            <div className={styles.data_types}>
              <div
                className={classNames(styles.choose_item, {
                  [styles.active]: activeItem === 0,
                })}
                onClick={() => handleItemClick(0)}
              >
                <p>Персональные данные</p>
              </div>
              <div
                className={classNames(styles.choose_item, {
                  [styles.active]: activeItem === 1,
                })}
                onClick={() => handleItemClick(1)}
              >
                <p>Карточка питомца</p>
              </div>
            </div>
            <div className={styles.data_main_info}>
              {activeItem === 0 ? (
                <>
                  <div className={styles.data_enter}>
                    <p>ФИО</p>
                    <input type="text" placeholder="ФИО" />
                  </div>
                  <div className={styles.data_enter}>
                    <p>Город</p>
                    <input type="text" placeholder="Город" />
                  </div>
                  <div className={styles.data_enter}>
                    <p>Телефон</p>
                    <input type="text" placeholder="Телефон" />
                  </div>
                  <div className={styles.data_enter}>
                    <p>Почта</p>
                    <input type="text" placeholder="Почта" />
                  </div>
                  <div className={styles.data_enter}>
                    <p>Telegram</p>
                    <input type="text" placeholder="Telegram" />
                  </div>
                </>
              ) : (
                <>
                  <div className={styles.data_enter}>
                    <p>Кличка</p>
                    <input type="text" placeholder="ФИО" />
                  </div>
                  <div className={styles.data_enter}>
                    <p>Питомец</p>
                    <input type="text" placeholder="ФИО" />
                  </div>
                  <div className={styles.data_enter}>
                    <p>Порода</p>
                    <input type="text" placeholder="ФИО" />
                  </div>
                  <div className={styles.data_enter}>
                    <p>Возраст</p>
                    <input type="text" placeholder="ФИО" />
                  </div>
                  <div className={styles.data_enter}>
                    <p>Тип крови</p>
                    <input type="text" placeholder="ФИО" />
                  </div>
                  <div className={styles.data_enter}>
                    <p>Вес</p>
                    <input type="text" placeholder="ФИО" />
                  </div>
                  <div className={styles.data_enter}>
                    <p>Действующие прививки</p>
                    <input type="text" placeholder="ФИО" />
                  </div>
                </>
              )}
            </div>
            {activeItem === 0 && (
              <div className={styles.show_contacts}>
                <p>Скрыть контакты</p>
              </div>
            )}
            <div className={styles.data_button}>Сохранить</div>
          </div>
          <div className={styles.data_right}>
            {activeItem === 0 && (
              <>
                <div className={styles.data_photo}>
                  <p>Фотография пользователя</p>
                  <div className={styles.photo_download}>
                    <img src="" alt="photo" />
                  </div>
                </div>
              </>
            )}
            {activeItem === 1 && (
              <>
                <div className={styles.data_photo}>
                  <p>Фотография питомца</p>
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
                  <button className="button">Заполнить форму</button>
                </div>
                <div className={styles.paginator}>
                  <p>Питомец 1</p>
                  <div className={styles.paginator_click}>
                    <div className={styles.paginator_left}>
                      <button>
                        <img src={LeftArrow} alt="left arrow" />
                      </button>
                    </div>
                    <div className={styles.paginator_right}>
                      <button>
                        <img src={RightArrow} alt="right arrow" />
                      </button>
                    </div>
                  </div>
                </div>
              </>
            )}
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
