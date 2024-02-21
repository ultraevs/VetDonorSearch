import React from "react";
import styles from "./Advice.module.css";

import DonorPerson from "../../assets/img/DonorPerson.svg";

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
                Lorem ipsum dolor sit amet consectetur. Proin malesuada commodo
                dignissim neque. Sed diam morbi leo blandit nisl mauris
                dignissim faucibus quis. Commodo vestibulum nulla nunc magnis.
                Lectus faucibus suspendisse auctor sed.
              </p>
            </div>
          </div>
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
                Lorem ipsum dolor sit amet consectetur. Proin malesuada commodo
                dignissim neque. Sed diam morbi leo blandit nisl mauris
                dignissim faucibus quis. Commodo vestibulum nulla nunc magnis.
                Lectus faucibus suspendisse auctor sed.
              </p>
            </div>
          </div>
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
                Lorem ipsum dolor sit amet consectetur. Proin malesuada commodo
                dignissim neque. Sed diam morbi leo blandit nisl mauris
                dignissim faucibus quis. Commodo vestibulum nulla nunc magnis.
                Lectus faucibus suspendisse auctor sed.
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Advice;
