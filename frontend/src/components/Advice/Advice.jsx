import React from "react";
import styles from "./Advice.module.css";

import DonorPerson from "../../assets/img/DonorPerson.svg";
import DonorPerson2 from "../../assets/img/DonorPerson2.svg";
import DonorPerson3 from "../../assets/img/DonorPerson3.svg";

const Advice = () => {
  return (
    <div className="container">
      <div className={styles.advice_block}>
        <h2 className="subtitle">Советы доноров</h2>
        <div className={styles.advice_wrapper}>
          <div className={styles.advice_item}>
            <div className={styles.item_person}>
              <div className={styles.person_photo}>
                <img src={DonorPerson} alt="DonorPerson" />
              </div>
              <div className={styles.person_info}>
                <div className={styles.person_info_name}>
                  <p>Глеб Таланцев</p>
                </div>
                <div className={styles.person_info_city}>
                  <p>Москва</p>
                </div>
              </div>
            </div>
            <div className={styles.item_text}>
              <p>
              Не все животные подходят для донорства крови. Убедитесь, что ваш питомец соответствует критериям, установленным вашей ветеринарной клиникой или организацией, которая принимает доноров.
              </p>
            </div>
          </div>
          <div className={styles.advice_item}>
            <div className={styles.item_person}>
              <div className={styles.person_photo}>
                <img src={DonorPerson2} alt="DonorPerson" />
              </div>
              <div className={styles.person_info}>
                <div className={styles.person_info_name}>
                  <p>Константин Цой</p>
                </div>
                <div className={styles.person_info_city}>
                  <p>Санкт-Петербург</p>
                </div>
              </div>
            </div>
            <div className={styles.item_text}>
              <p>
              Перед каждым сеансом донорства убедитесь, что ваш питомец находится в хорошем состоянии здоровья. Это включает в себя регулярные посещения ветеринара, контроль за питанием и уровнем активности.
              </p>
            </div>
          </div>
          <div className={styles.advice_item}>
            <div className={styles.item_person}>
              <div className={styles.person_photo}>
                <img src={DonorPerson3} alt="DonorPerson" />
              </div>
              <div className={styles.person_info}>
                <div className={styles.person_info_name}>
                  <p>Иван Лобода</p>
                </div>
                <div className={styles.person_info_city}>
                  <p>Челябинск</p>
                </div>
              </div>
            </div>
            <div className={styles.item_text}>
              <p>
              Если ваш питомец становится донором крови, обязательно придерживайтесь рекомендаций ветеринара относительно регулярности донорских сеансов.
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Advice;
