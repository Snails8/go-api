import React from "react";

import styles from './Sidebar.module.css'

export const Sidebar: React.FC = ({

}) => { 
    return (
        <div className={`${styles.base}`}>
            <div className={ `${styles.sidebar_item}` }>
                <div className={ `${styles.sidebar_text}`} >text</div>
            </div>
            <div className={ `${styles.sidebar_item}` }>
                <div className={ `${styles.sidebar_text}`} >text</div>
            </div>
            <div className={ `${styles.sidebar_item}` }>
                <div className={ `${styles.sidebar_text}`} >text</div>
            </div>
            <div className={ `${styles.sidebar_item}` }>
                <div className={ `${styles.sidebar_text}`} >text</div>
            </div>
            <div className={ `${styles.sidebar_item}` }>
                <div className={ `${styles.sidebar_text}`} >text</div>
            </div>
        </div>
    )

}