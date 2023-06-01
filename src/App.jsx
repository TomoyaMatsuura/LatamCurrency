import { Graph } from "./components/Graph";
import { ChakraProvider, HStack, Wrap, WrapItem } from "@chakra-ui/react";
//import axios from "axios";

import theme from "./theme/theme";

import { MBR } from "./components/affiliates/MBR";
import { MMX } from "./components/affiliates/MMX";
import { MCL } from "./components/affiliates/MCL";
import { MAR } from "./components/affiliates/MAR";
import { MPE } from "./components/affiliates/MPE";
import { MCO } from "./components/affiliates/MCO";

import { Header } from "./components/Header";
import { TableAffiliates } from "./components/TableAffiliates";
import { PutExcelButton } from "./components/PutExcelButton";

export default function App() {
  const affiliates = [MBR, MMX, MCL, MAR, MPE, MCO];

  return (
    <ChakraProvider theme={theme}>
      <Header />
      <Wrap pt={4}>
        <HStack>
          <TableAffiliates />
          <PutExcelButton affiliates={[1, 2, 3]} />
        </HStack>
      </Wrap>
      <Wrap p={4}>
        {affiliates.map((a) => (
          <WrapItem key={a.branch}>
            <Graph
              branch={a.branch}
              currency={a.currency}
              color={a.color}
              rateThisYear={a.rateThisYear}
              rateLastYear={a.rateLastYear}
              flag={a.flag}
              FinData={a.FinData}
            />
          </WrapItem>
        ))}
      </Wrap>
    </ChakraProvider>
  );
}
