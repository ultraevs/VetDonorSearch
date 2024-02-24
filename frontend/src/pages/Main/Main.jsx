import React, { useState, useEffect } from "react";
import styles from "./Main.module.css";

import Cookies from "js-cookie";

import Advice from "../../components/Advice/Advice";
import NeedDonor from "../../components/NeedDonor/NeedDonor";
import Hero from "../../components/Hero/Hero";
import MainLayout from "../../components/MainLayout/MainLayout";
import NeedToClinic from "../../components/NeedToClinic/NeedToClinic";
import Events from "../../components/Events/Events";
import TopDonors from "../../components/TopDonors/TopDonors";


// import axios from "axios";

const Main = () => {
  console.log(Cookies.get("Authorization"));

  // useEffect(() => {
  //   axios
  //     .get("https://vetdonor.shmyaks.ru/v1/profile", {
  //       withCredentials: true,
  //       headers: {
  //         Authorization:  `${Cookies.get("Authorization")}`
  //       },
  //     })
  //     .then(function (response) {
  //       console.log(response.data);
  //       setTitle(response.data)
  //     })
  //     .catch(function (error) {
  //       console.log(error);
  //     })
  //     .finally(function () {
  //     });
  // }, []);

  return (
    <>
      <MainLayout>
        <Hero />
        <Advice />
        <NeedDonor />
        <NeedToClinic />
        <Events />
        <TopDonors />
      </MainLayout>
    </>
  );
};

export default Main;
