import { Admin, Resource, ListGuesser } from "react-admin";
import simpleRestProvider from "ra-data-simple-rest";

const dataProvider = simpleRestProvider("http://localhost/api");

const App = () => (
  <Admin dataProvider={dataProvider}>
    <Resource name="ping" list={ListGuesser} />
  </Admin>
);

export default App;
