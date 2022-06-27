import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";

import { SamplePage2 } from "./pages/pageB";
import { Dashboard } from "./pages/dashboard/Dashboard";
import { SamplePage1 } from "./pages/pageA";
import { NotFound } from "./pages/notFound";
import { SamplePage3 } from "./pages/pageC";
import { SamplePage4, SamplePage4Child1, SamplePage4Child2, SamplePage4Child3 } from "./pages/pageD";


export const RouterConfig:React.VFC =() => {
  return (
    <>
     <BrowserRouter>
      <Routes>
        <Route index element={<Dashboard />} />
        <Route path="page1" element={<SamplePage1 />} />
        <Route path="page2" element={<SamplePage2 />} />
        <Route path="page3" element={<SamplePage3 Message="Hello Router" />} />
        <Route path="page4" element={<SamplePage4 />} >
          <Route index />
          <Route path="child1" element={<SamplePage4Child1 />} />
          <Route path="child2" element={<SamplePage4Child2 />} />
          <Route path=":cildid" element={<SamplePage4Child3 />} />
        </Route>
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
    </>
  );
}