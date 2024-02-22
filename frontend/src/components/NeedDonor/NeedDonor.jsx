import React from "react";
import styles from "./NeedDonor.module.css";

import Blood from "../../assets/img/blood.svg";
import Location from "../../assets/img/map.svg";
import Dog from "../../assets/img/dog.svg";
import PersonDog from "../../assets/img/personDog.svg";

const NeedDonor = () => {
  return (
    <div className="container">
      <div className={styles.needDonor_block}>
        <h2 className="subtitle">Срочно нужны доноры</h2>
        <div className={styles.needDonor_wrapper}>
          <div className={styles.needDonor_item}>
            <div className={styles.item_main}>
              <div className={styles.item_type}>
                <div className={styles.item_type_img}>
                  <img src={Blood} alt="type" />
                </div>
                <div className={styles.item_type_text}>
                  <p>Цельная кровь</p>
                  <p>DEA 1.1</p>
                </div>
              </div>
              <div className={styles.item_photo}>
                <img src={Dog} alt="dog" />
              </div>
              <div className={styles.item_info}>
                <div className={styles.item_info_name}>
                  <p>Бобик</p>
                  <p>Золотистый ретривер</p>
                </div>
                <div className={styles.item_info_adress}>
                  <div className={styles.item_info_adress_img}>
                    <img src={Location} alt="location" />
                  </div>
                  <div className={styles.item_info_adress_text}>
                    <p>г. Москва, ул. Мисисовская, д. 2, к. 2</p>
                  </div>
                </div>
                <hr />
                <div className={styles.item_info_found}>
                  <p>Ищет</p>
                  <div className={styles.item_info_found_person}>
                    <div className={styles.item_info_found_person_img}>
                      <img src={PersonDog} alt="person" />
                    </div>
                    <div className={styles.item_info_found_person_text}>
                      <p>Миша Евсеев</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div className={styles.item_date}>
              <p>26</p>
              <p>Фев</p>
            </div>
          </div>
          <div className={styles.needDonor_item}>
            <div className={styles.item_main}>
              <div className={styles.item_type}>
                <div className={styles.item_type_img}>
                  <img src={Blood} alt="type" />
                </div>
                <div className={styles.item_type_text}>
                  <p>Цельная кровь</p>
                  <p>DEA 1.1</p>
                </div>
              </div>
              <div className={styles.item_photo}>
                <img src={Dog} alt="dog" />
              </div>
              <div className={styles.item_info}>
                <div className={styles.item_info_name}>
                  <p>Бобик</p>
                  <p>Золотистый ретривер</p>
                </div>
                <div className={styles.item_info_adress}>
                  <div className={styles.item_info_adress_img}>
                    <img src={Location} alt="location" />
                  </div>
                  <div className={styles.item_info_adress_text}>
                    <p>г. Москва, ул. Мисисовская, д. 2, к. 2</p>
                  </div>
                </div>
                <hr />
                <div className={styles.item_info_found}>
                  <p>Ищет</p>
                  <div className={styles.item_info_found_person}>
                    <div className={styles.item_info_found_person_img}>
                      <img src={PersonDog} alt="person" />
                    </div>
                    <div className={styles.item_info_found_person_text}>
                      <p>Миша Евсеев</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div className={styles.item_date}>
              <p>26</p>
              <p>Фев</p>
            </div>
          </div>
          <div className={styles.needDonor_item}>
            <div className={styles.item_main}>
              <div className={styles.item_type}>
                <div className={styles.item_type_img}>
                  <img src={Blood} alt="type" />
                </div>
                <div className={styles.item_type_text}>
                  <p>Цельная кровь</p>
                  <p>DEA 1.1</p>
                </div>
              </div>
              <div className={styles.item_photo}>
                <img src={Dog} alt="dog" />
              </div>
              <div className={styles.item_info}>
                <div className={styles.item_info_name}>
                  <p>Бобик</p>
                  <p>Золотистый ретривер</p>
                </div>
                <div className={styles.item_info_adress}>
                  <div className={styles.item_info_adress_img}>
                    <img src={Location} alt="location" />
                  </div>
                  <div className={styles.item_info_adress_text}>
                    <p>г. Москва, ул. Мисисовская, д. 2, к. 2</p>
                  </div>
                </div>
                <hr />
                <div className={styles.item_info_found}>
                  <p>Ищет</p>
                  <div className={styles.item_info_found_person}>
                    <div className={styles.item_info_found_person_img}>
                      <img src={PersonDog} alt="person" />
                    </div>
                    <div className={styles.item_info_found_person_text}>
                      <p>Миша Евсеев</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div className={styles.item_date}>
              <p>26</p>
              <p>Фев</p>
            </div>
          </div>
          <div className={styles.needDonor_item}>
            <div className={styles.item_main}>
              <div className={styles.item_type}>
                <div className={styles.item_type_img}>
                  <img src={Blood} alt="type" />
                </div>
                <div className={styles.item_type_text}>
                  <p>Цельная кровь</p>
                  <p>DEA 1.1</p>
                </div>
              </div>
              <div className={styles.item_photo}>
                <img src={Dog} alt="dog" />
              </div>
              <div className={styles.item_info}>
                <div className={styles.item_info_name}>
                  <p>Бобик</p>
                  <p>Золотистый ретривер</p>
                </div>
                <div className={styles.item_info_adress}>
                  <div className={styles.item_info_adress_img}>
                    <img src={Location} alt="location" />
                  </div>
                  <div className={styles.item_info_adress_text}>
                    <p>г. Москва, ул. Мисисовская, д. 2, к. 2</p>
                  </div>
                </div>
                <hr />
                <div className={styles.item_info_found}>
                  <p>Ищет</p>
                  <div className={styles.item_info_found_person}>
                    <div className={styles.item_info_found_person_img}>
                      <img src={PersonDog} alt="person" />
                    </div>
                    <div className={styles.item_info_found_person_text}>
                      <p>Миша Евсеев</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div className={styles.item_date}>
              <p>26</p>
              <p>фев</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default NeedDonor;
