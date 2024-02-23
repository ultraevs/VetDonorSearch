import React from "react";
import { NavLink } from "react-router-dom"
import styles from "./Hero.module.css";

import HeroPic from "../../assets/img/hero.svg";
import Arrow from "../../assets/img/arrow.svg";
import WhiteArrow from "../../assets/img/WhiteArrow.svg";

const Hero = () => {
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

        <NavLink className={styles.hero_link} to="/Profile">
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
        </NavLink>

      </div>
      <div className={styles.bluebg_zone}></div>
    </div>
  );
};

export default Hero;
