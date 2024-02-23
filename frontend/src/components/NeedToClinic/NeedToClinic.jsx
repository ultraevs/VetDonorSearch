import React from "react";
import styles from "./NeedToClinic.module.css";

import { Link } from "react-router-dom";

import Geo from "../../assets/img/map.svg";
import PersonComeIn from "../../assets/img/PersonComeIn.svg";

const NeedToClinic = () => {
  return (
    <div className="container">
      <div className={styles.needToClinic_block}>
        <h2 className="subtitle">Потребности центров крови</h2>
        <div className={styles.need_wrapper}>
          <div className={styles.need_item}>
            <div className={styles.item_info}>
              <p>Название клиники</p>
              <div className={styles.item_info_geo}>
                <div className={styles.item_info_ego_img}>
                  <img src={Geo} alt="geo" />
                </div>
                <p>Ленинский проспект, 4</p>
              </div>
            </div>
            <div className={styles.item_info_types}>
              <div className={styles.types_wrapper}>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
              </div>
              <div className={styles.types_descriprion}>
                <div className={styles.types_true}>
                  <svg
                    width="12"
                    height="12"
                    viewBox="0 0 12 12"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <circle cx="6" cy="6" r="6" fill="#39C47D" />
                  </svg>{" "}
                  <span>Требуется</span>
                </div>
                <div className={styles.types_false}>
                  <svg
                    width="12"
                    height="12"
                    viewBox="0 0 12 12"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <circle cx="6" cy="6" r="6" fill="#C43939" />
                  </svg>{" "}
                  <span>Достаточно</span>
                </div>
              </div>
            </div>
            <div className={styles.info_comeIn}>
              <Link>
                <button className="button">Присоединиться</button>
              </Link>
              <div className={styles.comeIn_img}>
                <img src={PersonComeIn} alt="personComeIn" />
                <img src={PersonComeIn} alt="personComeIn" />
                <img src={PersonComeIn} alt="personComeIn" />
              </div>
            </div>
          </div>
          <div className={styles.need_item}>
            <div className={styles.item_info}>
              <p>Название клиники</p>
              <div className={styles.item_info_geo}>
                <div className={styles.item_info_ego_img}>
                  <img src={Geo} alt="geo" />
                </div>
                <p>Ленинский проспект, 4</p>
              </div>
            </div>
            <div className={styles.item_info_types}>
              <div className={styles.types_wrapper}>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
              </div>
              <div className={styles.types_descriprion}>
                <div className={styles.types_true}>
                  <svg
                    width="12"
                    height="12"
                    viewBox="0 0 12 12"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <circle cx="6" cy="6" r="6" fill="#39C47D" />
                  </svg>{" "}
                  <span>Требуется</span>
                </div>
                <div className={styles.types_false}>
                  <svg
                    width="12"
                    height="12"
                    viewBox="0 0 12 12"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <circle cx="6" cy="6" r="6" fill="#C43939" />
                  </svg>{" "}
                  <span>Достаточно</span>
                </div>
              </div>
            </div>
            <div className={styles.info_comeIn}>
              <Link>
                <button className="button">Присоединиться</button>
              </Link>
              <div className={styles.comeIn_img}>
                <img src={PersonComeIn} alt="personComeIn" />
                <img src={PersonComeIn} alt="personComeIn" />
                <img src={PersonComeIn} alt="personComeIn" />
              </div>
            </div>
          </div>
          <div className={styles.need_item}>
            <div className={styles.item_info}>
              <p>Название клиники</p>
              <div className={styles.item_info_geo}>
                <div className={styles.item_info_ego_img}>
                  <img src={Geo} alt="geo" />
                </div>
                <p>Ленинский проспект, 4</p>
              </div>
            </div>
            <div className={styles.item_info_types}>
              <div className={styles.types_wrapper}>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
                <div className={styles.types_item}>
                  <p>AB+</p>
                </div>
              </div>
              <div className={styles.types_descriprion}>
                <div className={styles.types_true}>
                  <svg
                    width="12"
                    height="12"
                    viewBox="0 0 12 12"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <circle cx="6" cy="6" r="6" fill="#39C47D" />
                  </svg>{" "}
                  <span>Требуется</span>
                </div>
                <div className={styles.types_false}>
                  <svg
                    width="12"
                    height="12"
                    viewBox="0 0 12 12"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <circle cx="6" cy="6" r="6" fill="#C43939" />
                  </svg>{" "}
                  <span>Достаточно</span>
                </div>
              </div>
            </div>
            <div className={styles.info_comeIn}>
              <Link>
                <button className="button">Присоединиться</button>
              </Link>
              <div className={styles.comeIn_img}>
                <img src={PersonComeIn} alt="personComeIn" />
                <img src={PersonComeIn} alt="personComeIn" />
                <img src={PersonComeIn} alt="personComeIn" />
              </div>
            </div>
          </div>
        </div>
        
      </div>
    </div>
  );
};

export default NeedToClinic;
