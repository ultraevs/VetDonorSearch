import React, { useState } from "react";
import styles from "./DropZone.module.css";

const DropZone = () => {
  function dragStartHandler(e) {
    e.preventDefault();
    setDrag(true);
  }

  function dragLeaveHandler(e) {
    e.preventDefault();
    setDrag(false);
  }

  function onDropHandler(e) {
    e.preventDefault();
    let files = [...e.dataTransfer.files];
    console.log(files);

    setDrag(false)
  }

  const [drag, setDrag] = useState(false);
  return (
    <div className={styles.app}>
      {drag ? (
        <div
          onDragStart={(e) => dragStartHandler(e)}
          onDragLeave={(e) => dragLeaveHandler(e)}
          onDragOver={(e) => dragStartHandler(e)}
          onDrop={(e) => onDropHandler(e)}
          className={styles.drop_area}
        >
          Отпустите файл, чтобы загрузить его
        </div>
      ) : (
        <div
          onDragStart={(e) => dragStartHandler(e)}
          onDragLeave={(e) => dragLeaveHandler(e)}
          onDragOver={(e) => dragStartHandler(e)}
          className={styles.drop}
        >
          Перетащите файл, чтобы загрузить его
        </div>
      )}
    </div>
  );
};

export default DropZone;