import axios from "axios";

export const createUser = async (name, mail, password) => {
  try {
    const response = await axios.post(
      "https://vetdonor.shmyaks.ru/v1/user_create",
      {
        email: `${mail}`,
        name: `${name}`,
        password: `${password}`,
      },
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
    return {success: true}
  } catch (error) {
    return {success: false, error: error};
  }
};

export const createClinic = async (name, address, mail, password) => {
  try {
    const response = await axios.post(
      "https://vetdonor.shmyaks.ru/v1/clinic_create",
      {
        name: `${name}`,
        address: `${address}`,
        email: `${mail}`,
        password: `${password}`,
      },
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
    return {success: true}
  } catch (error) {
    return {success: false, error: error};
  }
};
