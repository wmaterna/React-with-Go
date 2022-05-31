import React, { useContext } from "react";
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom';
import Navbar from "./components/Navbar";
import HomePage from "./screens/HomePage";
import ShopPage from "./screens/ShopPage";
import PaymentsPage from "./screens/PaymentsPage";
import CardPage from "./screens/CardPage";


function App() {
  return (
    <Router>
      <Navbar />
      <Routes path="/" >
          <Route path="/" element={<HomePage/>} />
          <Route path="/shop" element={<ShopPage/>} />
          <Route path="/payments" element={<PaymentsPage/>} />
          <Route path="/card" element={<CardPage/>} />
      </Routes>
    </Router>
  );
}

export default App;
