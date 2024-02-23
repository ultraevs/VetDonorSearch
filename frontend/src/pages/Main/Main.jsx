import React from "react";
import styles from "./Main.module.css";

import Advice from "../../components/Advice/Advice";
import NeedDonor from "../../components/NeedDonor/NeedDonor";
import Hero from "../../components/Hero/Hero";
import MainLayout from "../../components/MainLayout/MainLayout";
import NeedToClinic from "../../components/NeedToClinic/NeedToClinic";
import Events from "../../components/Events/Events";
import TopDonors from "../../components/TopDonors/TopDonors";

const Main = () => {
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
