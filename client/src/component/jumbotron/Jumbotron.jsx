import React from "react";
import bgImg from "../../assets/Jumbotron.png";
import iconbtn from "../../assets/iconbtn.png";
import { Link } from "react-router-dom";

function Jumbotron() {
  return (
    <>
      <div
        style={{
          marginTop: "10vh",
          backgroundImage: `url(${bgImg})`,
          height: "58vh",
          width: "auto",
          backgroundPosition: "center top",
          backgroundSize: "100%",
          backgroundRepeat: "no-repeat",
        }}
      >
        <Link to="/patient/consultation">
          <button
            style={{
              // margin: "12rem",
              // padding: "2rem",
              border: "0px",
              backgroundColor: "white",
              height: "55px",
              marginTop: "15.5rem",
              marginLeft: "11.5rem",
              backgroundImage: `url(${iconbtn})`,
              backgroundRepeat: "no-repeat",
              backgroundSize: "15%",
              // backgroundPosition: "left center ",
              backgroundPositionX: "10px",
              backgroundPositionY: "center",
            }}
          >
            <span
              style={{
                fontWeight: "700",
                color: "#ff6185",
                fontSize: "20px",
                marginLeft: "4rem",
              }}
              className="py-5"
            >
              Konsultasi Dengan Dokter
            </span>
          </button>
        </Link>
      </div>
    </>
  );
}
export default Jumbotron;
