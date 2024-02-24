import React, { useState } from "react";
import styles from "./ProfileInfo.module.css";

import classNames from "class-names";

import TelegramCircle from "../../assets/img/telegramCircle.svg";
import infoCat from "../../assets/img/infoCat.svg";
import infoGift from "../../assets/img/infoGift.svg";

import RightArrow from "../../assets/img/RightArrow.svg";
import LeftArrow from "../../assets/img/LeftArrow.svg";
import DropZone from "./components/DropZone/DropZone";

import useProfileData from "../../pages/Main/http";
import { sendInfoUser } from "./http";

const ProfileInfo = () => {
  const profileData = useProfileData();
  console.log(profileData);
  const [activeItem, setActiveItem] = useState(0);

  // User form
  const [inputName, setInputName] = useState("");
  const [inputCity, setInputCity] = useState("");
  const [inputNumber, setInputNumber] = useState("");
  const [inputMail, setInputMail] = useState("");
  const [inputTelegram, setInputTelegram] = useState("");

  const clearInputsUser = () => {
    setInputName(profileData?.name);
    setInputCity(profileData?.city);
    setInputNumber(profileData?.number);
    setInputMail(profileData?.mail);
    setInputTelegram(profileData?.telegram);
  };

  // Pet form
  const [inputPetName, setInputPetName] = useState("");
  const [inputPetType, setInputPetType] = useState("");
  const [inputPetGen, setInputPetGen] = useState("");
  const [inputPetAge, setInputPetAge] = useState("");
  const [inputPetTypeBlood, setInputPetTypeBlood] = useState("");
  const [inputPetWeight, setInputPetWeight] = useState("");
  const [inputPetArray, setInputPetArray] = useState("");

  const clearInputsPet = () => {
    setInputPetName("");
    setInputPetType("");
    setInputPetGen("");
    setInputPetAge("");
    setInputPetTypeBlood("");
    setInputPetWeight("");
    setInputPetArray("");
  };

  const userSaveInfo = async (city, email, name, path, phone, telegram) => {
    try {
      const response = await sendInfoUser(city, email, name, path, phone, telegram);
      console.log(response)
      if (response.success) {
      } else {
        alert(response.error)
      }
    } catch (error) {
      console.error(error);
    }
    clearInputsUser();
  };

  const petSaveInfo = () => {
    clearInputsPet();
  };

  const handleItemClick = (item) => {
    setActiveItem(item);
    if (item === 0) {
      clearInputsPet();
    } else {
      clearInputsUser();
    }
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
                    <input
                      type="text"
                      placeholder={`${profileData?.name}`}
                      value={inputName}
                      onChange={(event) => setInputName(event.target.value)}
                    />
                  </div>
                  <div className={styles.data_enter}>
                    <p>Город</p>
                    <input
                      type="text"
                      placeholder={`${profileData?.city}`}
                      value={inputCity}
                      onChange={(event) => setInputCity(event.target.value)}
                    />
                  </div>
                  <div className={styles.data_enter}>
                    <p>Телефон</p>
                    <input
                      type="text"
                      placeholder={`${profileData?.number}`}
                      value={inputNumber}
                      onChange={(event) => setInputNumber(event.target.value)}
                    />
                  </div>
                  <div className={styles.data_enter}>
                    <p>Почта</p>
                    <input
                      type="text"
                      placeholder={`${profileData?.email}`}
                      value={inputMail}
                      onChange={(event) => setInputMail(event.target.value)}
                    />
                  </div>
                  <div className={styles.data_enter}>
                    <p>Telegram</p>
                    <input
                      type="text"
                      placeholder={`${profileData?.telegram}`}
                      value={inputTelegram}
                      onChange={(event) => setInputTelegram(event.target.value)}
                    />
                  </div>
                </>
              ) : (
                <>
                  <div className={styles.data_enter}>
                    <p>Кличка</p>
                    <input
                      type="text"
                      placeholder="Кличка"
                      value={inputPetName}
                      onChange={(event) => setInputPetName(event.target.value)}
                    />
                  </div>
                  <div className={styles.data_enter}>
                    <p>Питомец</p>
                    <input
                      type="text"
                      placeholder="Питомец"
                      value={inputPetType}
                      onChange={(event) => setInputPetType(event.target.value)}
                    />
                  </div>
                  <div className={styles.data_enter}>
                    <p>Порода</p>
                    <input
                      type="text"
                      placeholder="Порода"
                      value={inputPetGen}
                      onChange={(event) => setInputPetGen(event.target.value)}
                    />
                  </div>
                  <div className={styles.data_enter}>
                    <p>Возраст</p>
                    <input
                      type="text"
                      placeholder="Возраст"
                      value={inputPetAge}
                      onChange={(event) => setInputPetAge(event.target.value)}
                    />
                  </div>
                  <div className={styles.data_enter}>
                    <p>Тип крови</p>
                    <input
                      type="text"
                      placeholder="Тип крови"
                      value={inputPetTypeBlood}
                      onChange={(event) =>
                        setInputPetTypeBlood(event.target.value)
                      }
                    />
                  </div>
                  <div className={styles.data_enter}>
                    <p>Вес</p>
                    <input
                      type="text"
                      placeholder="Вес"
                      value={inputPetWeight}
                      onChange={(event) =>
                        setInputPetWeight(event.target.value)
                      }
                    />
                  </div>
                  <div className={styles.data_enter}>
                    <p>Действующие прививки</p>
                    <input
                      type="text"
                      placeholder="Действующие прививки"
                      value={inputPetArray}
                      onChange={(event) => setInputPetArray(event.target.value)}
                    />
                  </div>
                </>
              )}
            </div>
            {activeItem === 0 && (
              <div className={styles.show_contacts}>
                <p>Скрыть контакты</p>
              </div>
            )}
            <div
              className={styles.data_button}
              onClick={() => {
                if (activeItem === 0) {
                  userSaveInfo(inputCity, inputMail, inputName, "path", inputNumber, inputTelegram);
                } else {
                  petSaveInfo();
                }
              }}
            >
              Сохранить
            </div>
          </div>
          <div className={styles.data_right}>
            {activeItem === 0 && (
              <>
                <div className={styles.data_photo}>
                  <p>Фотография пользователя</p>
                  <DropZone />
                </div>
              </>
            )}
            {activeItem === 1 && (
              <>
                <div className={styles.data_photo}>
                  <p>Фотография питомца</p>
                  <DropZone />
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