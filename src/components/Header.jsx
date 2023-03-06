import { Flex, Heading } from "@chakra-ui/react";
import { memo } from "react";
import { MenuNote } from "./MenuNote";

export const Header = memo(() => {
  return (
    <Flex
      as="nav"
      bg="teal.500"
      color="white"
      align="center"
      justify="space-between"
    >
      <Heading as="h1" textAlign="center" padding="2">
        中南米為替レート
      </Heading>
      <MenuNote />
    </Flex>
  );
});
