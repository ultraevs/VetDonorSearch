import React from "react";
import { YMaps, Map, Placemark, GeolocationControl } from "@pbe/react-yandex-maps";

const Maps = () => {
  return (
    <div className="container">
      <YMaps>
        <Map
          style={{ marginTop: 42, borderRadius: 12, overflow: "hidden", width: "100%", height: 600 }}
          defaultState={{
            center: [59.938678, 30.314474],
            zoom: 9,
            controls: ["zoomControl", "fullscreenControl"],
          }}
          modules={["control.ZoomControl", "control.FullscreenControl"]}
        >
          <GeolocationControl options={{ float: "left" }} />
          <Placemark
            modules={["geoObject.addon.balloon"]}
            defaultGeometry={[59.898319, 30.307682]}
            properties={{
              balloonContentHeader:
                "Red Вет",
              balloonContentBody:
                "10:00-21:00",
              balloonBottom: true,
            }}
          />
          <Placemark
          modules={["geoObject.addon.balloon"]}
          defaultGeometry={[59.924330, 30.279091]}
          properties={{
            balloonContentHeader:
            "Вместе",
            balloonContentBody:
              "10:00-20:00",
          }}
        />
        <Placemark
          modules={["geoObject.addon.balloon"]}
          defaultGeometry={[59.900822, 30.456492]}
          properties={{
            balloonContentHeader:
            "Свой доктор",
            balloonContentBody:
              "Круглосуточно",
          }}
        />
        <Placemark
          modules={["geoObject.addon.balloon"]}
          defaultGeometry={[59.886869, 30.318417]}
          properties={{
            balloonContentHeader:
            "БАРСЕЛЬ",
            balloonContentBody:
              "10:00-22:00",
          }}
        />
        <Placemark
          modules={["geoObject.addon.balloon"]}
          defaultGeometry={[59.913465, 30.327402]}
          properties={{
            balloonContentHeader:
            "Ветеринарная клиника доктора Бушаровой",
            balloonContentBody:
              "11:00-21:00",
          }}
        />
        <Placemark
          modules={["geoObject.addon.balloon"]}
          defaultGeometry={[59.940532, 30.265803]}
          properties={{
            balloonContentHeader:
            "Ветлекаръ",
            balloonContentBody:
              "Круглосуточно",
          }}
        />
        <Placemark
          modules={["geoObject.addon.balloon"]}
          defaultGeometry={[59.963747, 30.364817]}
          properties={{
            balloonContentHeader:
            "Прайд",
            balloonContentBody:
              "Круглосуточно",
          }}
        />
        <Placemark
          modules={["geoObject.addon.balloon"]}
          defaultGeometry={[59.943915, 30.479324]}
          properties={{
            balloonContentHeader:
            "Вега",
            balloonContentBody:
              "Круглосуточно",
          }}
        />
        </Map>
      </YMaps>
    </div>
  );
};

export default Maps;