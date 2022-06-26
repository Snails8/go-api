import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";

import { SamplePage2 } from "./pages/pageB";
import { Dashboard } from "./pages/dashboard/Dashboard";
import { SamplePage1 } from "./pages/pageA";


export const RouterConfig:React.VFC =() => {
  return (
    <>
     <BrowserRouter>
      <Routes>
        <Route index element={<Dashboard />} />
        <Route path="page1" element={<SamplePage1 />} />
        <Route path="page2" element={<SamplePage2 />} />
      </Routes>
    </BrowserRouter>
    </>
  );
}