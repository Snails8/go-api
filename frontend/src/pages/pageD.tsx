import React, { FC } from "react";
import { Outlet, Link, useNavigate} from "react-router-dom";

export const SamplePage4:FC = () => {
  const navigate = useNavigate();

  return (
    <>
      <h3>Sample Page 4</h3>
      <ul>
        <li><Link to="child1">Show Child1</Link></li>
        <li><Link to="child2">Show Child2</Link></li>
      </ul>
      <button onClick={()=> navigate("") }>clear</button>
      <Outlet  />
    </>
  );
}

export const SamplePage4Child1:FC = () => {
  return <h3>Sample Page 4 Child1</h3>;
}
export const SamplePage4Child2:FC = () => {
  return <h3>Sample Page 4 Child2</h3>;
}