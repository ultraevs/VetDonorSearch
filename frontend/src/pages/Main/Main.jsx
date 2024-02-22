import React from "react";
import styles from "./Main.module.css";

import Advice from "../../components/Advice/Advice";
import NeedDonor from "../../components/NeedDonor/NeedDonor";
import Hero from "../../components/Hero/Hero";
import MainLayout from "../../components/MainLayout/MainLayout";

const Main = () => {
  return (
    <>
      <MainLayout>
        <Hero />
        <Advice />
        <NeedDonor />
      </MainLayout>
    </>
  );
};

export default Main;
