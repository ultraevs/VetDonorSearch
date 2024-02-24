import { useState, useEffect } from 'react';
import axios from 'axios';
import Cookies from "js-cookie";

const useProfileData = () => {
  const [profileData, setProfileData] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get("https://vetdonor.shmyaks.ru/v1/profile", {
          withCredentials: true,
          headers: {
            Authorization:  `${Cookies.get("Authorization")}`
          },
        });
        setProfileData(response.data);
      } catch (error) {
        console.log(error);
      }
    };

    fetchData();
  }, []);

  return profileData;
};

export default useProfileData;