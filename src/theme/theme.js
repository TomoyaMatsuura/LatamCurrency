import { extendTheme } from "@chakra-ui/react";

const theme = extendTheme({
  styles: {
    global: {
      body: {
        backgroundColor: "blue.50",
        color: "gray.800"
      }
    }
  }
});
export default theme;
