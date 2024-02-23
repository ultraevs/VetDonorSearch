import React from "react";
import styles from "./Footer.module.css";

import Telegram from "../../assets/img/telegram.svg";
import Mail from "../../assets/img/mail.svg";

const Footer = () => {
  return (
    <div className={styles.footer_bg}>
      <div className={styles.footer_block}>
        <div className={styles.footer_items}>
          <div className={styles.footer_item}>
            <div className={styles.item_telegram}>
              <img src={Telegram} alt="telegram" />
              <a href="https://t.me/vetdonor">@vetdonor</a>
            </div>
            <div className={styles.item_mail}>
              <img src={Mail} alt="mail" />
              <a href="mailto:smyakneksbimisis@gmail.com">
                smyakneksbimisis@gmail.com
              </a>
            </div>
          </div>
          <hr />
          <div className={styles.footer_item}>
            <div className={styles.item_brend}>
              <p>© VetDonor</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Footer;
