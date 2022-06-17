import '../src/assets/app.scss';
import {useEffect} from "react";
import { useSetRecoilState, RecoilRoot} from "recoil";
import {loginUserState} from "../src/states/loginUser";
import * as loginUser from "../src/lib/loginUser";

const MyApp = ({ Component, pageProps }) => {

  return (
      <>
        <RecoilRoot>
          <AppInit />
          <Component {...pageProps} />
        </RecoilRoot>
      </>
  )
}

function AppInit() {
  const setLoginUser = useSetRecoilState(loginUserState)

  // Middleware
  useEffect(() => {
    (async () => {
      try {
        const user = await loginUser.fetchLoginUser();
        setLoginUser(user);
      } catch(e) {
        setLoginUser(null);
      }
    })();
  }, [])

  return null
}

export default MyApp
