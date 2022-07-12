import React from "react";
import { Link } from "react-router-dom";
import { SidebarBg } from "../../atoms/Sidebar/Sidebar";

import styles from './Sidebar.module.css'

export const Sidebar: React.FC = ({

}) => { 
    return (
        <SidebarBg>
            <Link to="page1">
                <div className={ `${styles.sidebar_item}` }>
                    <div className={ `${styles.sidebar_text}`} >page1</div>
                </div>
            </Link>
            <Link to="page2">
                <div className={ `${styles.sidebar_item}` }>
                    <div className={ `${styles.sidebar_text}`} >page2</div>
                </div>
            </Link>
            <Link to="page3">
                <div className={ `${styles.sidebar_item}` }>
                    <div className={ `${styles.sidebar_text}`} >page3</div>
                </div>
            </Link>
            <Link to="page4">
                <div className={ `${styles.sidebar_item}` }>
                    <div className={ `${styles.sidebar_text}`} >page4</div>
                </div>
            </Link>
            <Link to="users">
                <div className={ `${styles.sidebar_item}` }>
                    <div className={ `${styles.sidebar_text}`} >ユーザー管理</div>
                </div>
            </Link>
        </SidebarBg>
    )
}