import React from "react";
import styles from "./Events.module.css";

import { Link } from "react-router-dom";

import MSHA from "../../assets/img/MSHA.svg";
import ITMO from "../../assets/img/ITMO.svg";
import MIET from "../../assets/img/MIET.svg";

const Events = () => {
  return (
    <div className="container">
      <div className={styles.events_block}>
        <h2 className="subtitle">Мероприятия</h2>
        <div className={styles.events_wrapper}>
          <div style={{ width: 497 }} className={styles.events_item}>
            <div className={styles.item_title}>
              <div className={styles.item_title}>
                <img src={MSHA} alt="Event" />
              </div>
              <div className={styles.item_nameing}>
                <p>
                  РГАУ-МСХА
                  <br />
                  им. К. А. Тимирязева
                </p>
              </div>
            </div>
            <div className={styles.item_description}>
              <p>Выездная донорская акция</p>
            </div>
            <div className={styles.item_connect}>
              <Link>
                <button className="button">Записаться</button>
              </Link>
            </div>
          </div>
          <div style={{ width: 393 }} className={styles.events_item}>
            <div className={styles.item_title}>
              <div className={styles.item_title}>
                <img src={ITMO} alt="Event" />
              </div>
              <div className={styles.item_nameing}>
                <p>
                  Добровольческий
                  <br />
                  центр ИТМО
                </p>
              </div>
            </div>
            <div className={styles.item_description}>
              <p>День Донора в ИТМО</p>
            </div>
            <div className={styles.item_connect}>
              <Link>
                <button className="button">Записаться</button>
              </Link>
            </div>
          </div>
          <div style={{ width: 290 }} className={styles.events_item}>
            <div className={styles.item_title}>
              <div className={styles.item_title}>
                <img src={MIET} alt="Event" />
              </div>
              <div className={styles.item_nameing}>
                <p>НИУ МИЭТ</p>
              </div>
            </div>
            <div className={styles.item_description}>
              <p>Донорское движение</p>
            </div>
            <div className={styles.item_connect}>
              <Link>
                <button className="button">Записаться</button>
              </Link>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Events;
