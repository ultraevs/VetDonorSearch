import React, { useState } from "react";
import { NavLink } from "react-router-dom"
import styles from "./Hero.module.css";

import HeroPic from "../../assets/img/hero.svg";
import Arrow from "../../assets/img/arrow.svg";
import WhiteArrow from "../../assets/img/WhiteArrow.svg";
import Modal from "../Modal/Modal";
import Close from './../../assets/icon/close.svg'

const Hero = () => {
  const [modalActive, setModalActive] = useState(false)

  return (
    <div className={styles.hero_block}>
      <div className={styles.hero_info}>
        <NavLink className={styles.hero_link} to="/Info">
          <div className={styles.hero_info_top}>
            <div className={styles.hero_text}>
              <h1>Донорство крови для животных</h1>
              <div className={styles.hero_next}>
                <div className={styles.hero_next_text}>
                  <p>
                    Узнайте о преимуществах и требованиях для того, чтобы ваш
                    питомец стал героем.
                  </p>
                </div>
                <div className={styles.arrow}>
                  <img src={Arrow} alt="arrow" />
                </div>
              </div>
            </div>
          </div>
        </NavLink>

        <button className={styles.hero_button} onClick={() => setModalActive(true)}>
          <div className={styles.hero_info_bottom}>
            <div className={styles.hero_text}>
              <h1>Запрос крови для животного</h1>
              <div className={styles.hero_next}>
                <div className={styles.hero_next_text}>
                  <p>
                    Получите неотложную помощь прямо сейчас, обратившись к нам.
                  </p>
                </div>
                <div className={styles.arrow}>
                  <img src={WhiteArrow} alt="arrow" />
                </div>
              </div>
            </div>
            <div className={styles.hero_photo}>
              <img src={HeroPic} alt="hero" />
            </div>
          </div>
        </button>

      </div>
      <div className={styles.bluebg_zone}></div>

      <Modal active={modalActive} setActive={setModalActive}>
        <button className={styles.close} onClick={() => setModalActive(false)}>
          <img src={Close} alt="" />
        </button>
        <form className={styles.forma}>
          <div className={styles.form_text}>
            <label >Причина</label>
            <input placeholder='Не указано' type="text" />
          </div>
          <div className={styles.form_text}>
            <label >Ветклиника</label>
            <input placeholder='Не указано' type="text" />
          </div>
          <div className={styles.form_text}>
            <label >Группа крови</label>
            <input placeholder='Не указано' type="text" />
          </div>
          <div className={styles.form_text}>
            <label >Количество крови (мл)</label>
            <input placeholder='Не указано' type="text" />
          </div>
          <div className={styles.form_text}>
            <label >Дата, до которой поиск актуален</label>
            <input placeholder='Не указано' type="text" />
          </div>
        </form>
        <div className={styles.post_btn} onClick={() => setModalActive(false)}>
          <button className={styles.post_form}>
            Отправить
          </button>
        </div>

      </Modal>
    </div>
  );
};

export default Hero;
