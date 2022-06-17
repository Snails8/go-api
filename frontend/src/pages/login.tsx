import Layout from "../src/components/_layouts/Layout";
import * as loginUser from "../src/lib/loginUser";
import {Main} from "../src/components/login/Main";


const LoginPage = () => {


  return (
      <>
        <Layout title="Login">
          <Main />
        </Layout>
      </>
  )
}

export default LoginPage;
