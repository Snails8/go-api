import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";

import { SamplePage2 } from "./pages/pageB";
import { Dashboard } from "./pages/dashboard/Dashboard";
import { SamplePage1 } from "./pages/pageA";
import { NotFound } from "./pages/notFound";
import { SamplePage3 } from "./pages/pageC";


export const RouterConfig:React.VFC =() => {
  return (
    <>
     <BrowserRouter>
      <Routes>
        <Route index element={<Dashboard />} />
        <Route path="page1" element={<SamplePage1 />} />
        <Route path="page2" element={<SamplePage2 />} />
        <Route path="page3" element={<SamplePage3 Message="Hello Router" />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
    </>
  );
}