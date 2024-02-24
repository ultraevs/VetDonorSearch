import React, { useState } from "react";
import styles from "./ProfileClinicInfo.module.css";

import TelegramCircle from "../../assets/img/telegramCircle.svg";

import useProfileData from "../../pages/Main/http";

const ProfileClinicInfo = () => {
  // const typesBlood = [
  //   { "AB+": false },
  //   { "AB+": false },
  //   { "CD+": false },
  //   { "AB+": false },
  //   { "AA+": false },
  //   { "AB+": true },
  //   { "BB+": false },
  //   { "AB+": false },
  //   { "AB+": false },
  //   { "AB+": false },
  // ];

  const arrayEx = ["AB+", "AA+", "BB+", "B+", "BC+", "CC+", "C+", "ABs+", "dB+"];

  const [activeArr, setActiveArr] = useState([])

  const handleClick = (item) => {
    if (activeArr.includes(item)) {
      setActiveArr([...activeArr.filter((elem) => elem !== item)])
    } else {
      setActiveArr([...activeArr, item])
    }
  };
  
  const profileData = useProfileData();
  return (
    <div className="container">
      <div className={styles.profileInfo_block}>
        <div className={styles.profile_data}>
          <div className={styles.data_left}>
            <div className={styles.data_main_info}>
              <div className={styles.data_enter}>
                <p>Название клиники</p>
                <input type="text" placeholder={`${profileData?.name}`} />
              </div>
              <div className={styles.data_enter}>
                <p>Адресс</p>
                <input type="text" placeholder={`${profileData?.adress}`} />
              </div>
              <div className={styles.data_enter}>
                <p>Телефон</p>
                <input type="text" placeholder={`${profileData?.number}`} />
              </div>
              <div className={styles.data_enter}>
                <p>Почта</p>
                <input type="text" placeholder={`${profileData?.email}`} />
              </div>
              <div className={styles.data_enter}>
                <p>Telegram</p>
                <input type="text" placeholder={`${profileData?.telegram}`} />
              </div>
            </div>

            <div className={styles.show_contacts}>
              <p>Скрыть контакты</p>
            </div>
            <div className={styles.data_button}>Сохранить</div>
          </div>
          <div className={styles.data_right}>
            <div className={styles.needable}>
              <p>
                Потребность в крови
                <br />
                <span>Отметьте группы крови, в которых нуждается центр</span>
              </p>
              <div className={styles.typesOfBlood}>
                <div className={styles.types_wrapper}>
                  {arrayEx.map((item) => {
                    return (
                      <div
                        key={item}
                        className={`${styles.types_item} ${activeArr.includes(item) ? `${styles.activeType}` : ``}`}
                        onClick={() => handleClick(item)}
                      >
                        <p>{item}</p>
                      </div>
                    );
                  })}
                </div>
              </div>
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
        </div>
      </div>
    </div>
  );
};

export default ProfileClinicInfo;
