import React from "react";
import Layout from "../../components/Layout/Layout";

import styles from "./Start.module.css"

import { Link } from "react-router-dom";

const Start = () => {
  return (
    <Layout>
      <div className={styles.main_info}>
        <h1>Помощь животным</h1>
        <p>
          Vetdonor hello помогает найти доноров
          <br />
          крови для ваших питомцев
        </p>
        <Link to="/auth">
          <button>Присоединиться</button>
        </Link>
      </div>
    </Layout>
  );
};

export default Start;
