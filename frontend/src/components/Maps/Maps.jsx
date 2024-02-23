import React from "react";
import { YMaps, Map, Placemark, GeolocationControl } from "@pbe/react-yandex-maps";

const Maps = () => {
  return (
    <div className="container">
      <YMaps>
        <Map
        style={{marginTop: 42, borderRadius: 12, overflow: "hidden", width: "100%", height: 600}}
          defaultState={{
            center: [55.75, 37.57],
            zoom: 9,
            controls: ["zoomControl", "fullscreenControl"],
          }}
          modules={["control.ZoomControl", "control.FullscreenControl"]}
        >
          <GeolocationControl options={{ float: "left" }} />
          <Placemark
            modules={["geoObject.addon.balloon"]}
            defaultGeometry={[55.75, 37.57]}
            // properties={{
            //   tipContent:
            //     "This is a block with information attached to the bottom",
            //   // tipContentBottom: true,
            // }}        
            properties={{
              balloonContentBody:
                "This is balloon loaded by the Yandex.Maps API module system",
              balloonBottom: true,
            }}  
          />
          <Placemark
            modules={["geoObject.addon.balloon"]}
            defaultGeometry={[53, 38]}
            properties={{
              balloonContentBody:
                "This is balloon loaded by the Yandex.Maps API module system",
            }}
          />
        </Map>
      </YMaps>
    </div>
  );
};

export default Maps;
