import React from "react";
import styles from "./TopDonors.module.css";

import Donor from "../../assets/img/Donor.svg"
import Yellow from "../../assets/img/YellowCard.svg"
import Purple from "../../assets/img/PurpleType.svg"
import Red from "../../assets/img/RedType.svg"

const TopDonors = () => {
  return (
    <div className="container">
      <div className={styles.topDonors_block}>
        <h2 className="subtitle">Топ доноров</h2>
        <div className={styles.topDonors_wrapper}>
          <div className={styles.topDonors_item}>
            <div className={styles.item_person}>
              <div className={styles.item_person_img}>
                <img src={Donor} alt="person" />
              </div>
              <div className={styles.item_person_name}>
                <p>Иван Лобода</p>
              </div>
            </div>
            <div className={styles.item_count}>
                <p>Донации: <span>9</span></p>
            </div>
            <div className={styles.item_types}>
                <div className={styles.types_type}>
                    <img src={Yellow} alt="first type" />
                    <span>2</span>
                </div>
                <div className={styles.types_type}>
                    <img src={Red} alt="second type" />
                    <span>3</span>
                </div>
                <div className={styles.types_type}>
                    <img src={Purple} alt="third type" />
                    <span>7</span>
                </div>
            </div>
          </div>
          <div className={styles.topDonors_item}>
            <div className={styles.item_person}>
              <div className={styles.item_person_img}>
                <img src={Donor} alt="person" />
              </div>
              <div className={styles.item_person_name}>
                <p>Иван Лобода</p>
              </div>
            </div>
            <div className={styles.item_count}>
                <p>Донации: <span>9</span></p>
            </div>
            <div className={styles.item_types}>
                <div className={styles.types_type}>
                    <img src={Yellow} alt="first type" />
                    <span>2</span>
                </div>
                <div className={styles.types_type}>
                    <img src={Red} alt="second type" />
                    <span>3</span>
                </div>
                <div className={styles.types_type}>
                    <img src={Purple} alt="third type" />
                    <span>7</span>
                </div>
            </div>
          </div>
          <div className={styles.topDonors_item}>
            <div className={styles.item_person}>
              <div className={styles.item_person_img}>
                <img src={Donor} alt="person" />
              </div>
              <div className={styles.item_person_name}>
                <p>Иван Лобода</p>
              </div>
            </div>
            <div className={styles.item_count}>
                <p>Донации: <span>9</span></p>
            </div>
            <div className={styles.item_types}>
                <div className={styles.types_type}>
                    <img src={Yellow} alt="first type" />
                    <span>2</span>
                </div>
                <div className={styles.types_type}>
                    <img src={Red} alt="second type" />
                    <span>3</span>
                </div>
                <div className={styles.types_type}>
                    <img src={Purple} alt="third type" />
                    <span>7</span>
                </div>
            </div>
          </div>
          <div className={styles.topDonors_item}>
            <div className={styles.item_person}>
              <div className={styles.item_person_img}>
                <img src={Donor} alt="person" />
              </div>
              <div className={styles.item_person_name}>
                <p>Иван Лобода</p>
              </div>
            </div>
            <div className={styles.item_count}>
                <p>Донации: <span>9</span></p>
            </div>
            <div className={styles.item_types}>
                <div className={styles.types_type}>
                    <img src={Yellow} alt="first type" />
                    <span>2</span>
                </div>
                <div className={styles.types_type}>
                    <img src={Red} alt="second type" />
                    <span>3</span>
                </div>
                <div className={styles.types_type}>
                    <img src={Purple} alt="third type" />
                    <span>7</span>
                </div>
            </div>
          </div>
          <div className={styles.topDonors_item}>
            <div className={styles.item_person}>
              <div className={styles.item_person_img}>
                <img src={Donor} alt="person" />
              </div>
              <div className={styles.item_person_name}>
                <p>Иван Лобода</p>
              </div>
            </div>
            <div className={styles.item_count}>
                <p>Донации: <span>9</span></p>
            </div>
            <div className={styles.item_types}>
                <div className={styles.types_type}>
                    <img src={Yellow} alt="first type" />
                    <span>2</span>
                </div>
                <div className={styles.types_type}>
                    <img src={Red} alt="second type" />
                    <span>3</span>
                </div>
                <div className={styles.types_type}>
                    <img src={Purple} alt="third type" />
                    <span>7</span>
                </div>
            </div>
          </div>
          <div className={styles.topDonors_item}>
            <div className={styles.item_person}>
              <div className={styles.item_person_img}>
                <img src={Donor} alt="person" />
              </div>
              <div className={styles.item_person_name}>
                <p>Иван Лобода</p>
              </div>
            </div>
            <div className={styles.item_count}>
                <p>Донации: <span>9</span></p>
            </div>
            <div className={styles.item_types}>
                <div className={styles.types_type}>
                    <img src={Yellow} alt="first type" />
                    <span>2</span>
                </div>
                <div className={styles.types_type}>
                    <img src={Red} alt="second type" />
                    <span>3</span>
                </div>
                <div className={styles.types_type}>
                    <img src={Purple} alt="third type" />
                    <span>7</span>
                </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default TopDonors;
