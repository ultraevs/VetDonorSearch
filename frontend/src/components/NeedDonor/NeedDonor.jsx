import React, { useEffect, useState } from "react";
import styles from "./NeedDonor.module.css";

import Blood from "../../assets/img/blood.svg";
import Location from "../../assets/img/map.svg";
import Cat from "../../assets/img/cat1.svg";
import Cat2 from "../../assets/img/cat2.svg";
import Dog from "../../assets/img/dog.svg";
import Dog2 from "../../assets/img/dog2.svg";

import PersonCat from "../../assets/img/personCat1.svg";
import PersonCat2 from "../../assets/img/personCat2.svg";
import PersonDog from "../../assets/img/personDog.svg";
import PersonDog2 from "../../assets/img/personDog2.svg";


const NeedDonor = () => {
  const [dogsCards, setDogsCards] = useState([]);
  const [loading, setLoading] = useState(false);

  // useEffect(() => {
  //   setLoading(true)
  //   axios
  //     .get("/user?ID=12345")
  //     .then(function (response) {
  //       setDogsCards(response.data);
  //     })
  //     .catch(function (error) {
  //       console.log(error);
  //     })
  //     .finally(function () {
  //       setLoading(false)
  //     });
  // }, []);

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
                  <p>A</p>
                </div>
              </div>
              <div className={styles.item_photo}>
                <img src={Cat} alt="cat" />
              </div>
              <div className={styles.item_info}>
                <div className={styles.item_info_name}>
                  <p>Барсик</p>
                  <p>Ли Хуа</p>
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
                      <p>Клиника Прайд</p>
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
                  <p>B</p>
                </div>
              </div>
              <div className={styles.item_photo}>
                <img src={Cat2} alt="cat" />
              </div>
              <div className={styles.item_info}>
                <div className={styles.item_info_name}>
                  <p>Джонни</p>
                  <p>Британец</p>
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
                      <img src={PersonCat} alt="person" />
                    </div>
                    <div className={styles.item_info_found_person_text}>
                      <p>Клиника Вместе</p>
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
                  <p>DEA 1.1+</p>
                </div>
              </div>
              <div className={styles.item_photo}>
                <img src={Dog} alt="dog" />
              </div>
              <div className={styles.item_info}>
                <div className={styles.item_info_name}>
                  <p>Мухтар</p>
                  <p>Немецкая овчарка</p>
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
                      <img src={PersonCat2} alt="person" />
                    </div>
                    <div className={styles.item_info_found_person_text}>
                      <p>Иван Лобода</p>
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
                  <p>DEA 1.1-</p>
                </div>
              </div>
              <div className={styles.item_photo}>
                <img src={Dog2} alt="dog" />
              </div>
              <div className={styles.item_info}>
                <div className={styles.item_info_name}>
                  <p>Шон</p>
                  <p>Алабай</p>
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
                      <img src={PersonDog2} alt="person" />
                    </div>
                    <div className={styles.item_info_found_person_text}>
                      <p>Михаил Евсеев</p>
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
