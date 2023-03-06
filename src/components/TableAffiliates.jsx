import {
  Table,
  Thead,
  Tbody,
  Tfoot,
  Tr,
  Th,
  Td,
  TableContainer,
  Stack
} from "@chakra-ui/react";

import { MBR } from "./affiliates/MBR";
import { MMX } from "./affiliates/MMX";
import { MCL } from "./affiliates/MCL";
import { MAR } from "./affiliates/MAR";
import { MPE } from "./affiliates/MPE";
import { MCO } from "./affiliates/MCO";

export const TableAffiliates = () => {
  const affiliates = [MBR, MMX, MCL, MAR, MPE, MCO];
  const month = [
    "Apr",
    "May",
    "Jun",
    "Jul",
    "Sep",
    "Oct",
    "Nov",
    "Dec",
    "Jan",
    "Feb",
    "Mar"
  ];
  return (
    <Stack pl={5} pr={5}>
      <TableContainer
        borderRadius="10px"
        shadow="md"
        p={2}
        m={2}
        bg="white"
        color="gray.500"
        overflowX="auto"
      >
        <Table variant="simple" size="sm">
          <Thead>
            <Tr>
              <Th>国</Th>
              <Th>通貨</Th>
              {month.map((m) => (
                <Th isNumeric>{m}</Th>
              ))}
            </Tr>
          </Thead>
          <Tbody>
            {affiliates.map((a) => (
              <Tr>
                <Td>{a.country}</Td>
                <Td>{a.currency}</Td>
                {a.rateThisYear.map((m) => (
                  <Td isNumeric>{m}</Td>
                ))}
              </Tr>
            ))}
          </Tbody>
          <Tfoot>
            {/* <Tr>
            <Th>To convert</Th>
            <Th>into</Th>
            <Th isNumeric>multiply by</Th>
          </Tr> */}
          </Tfoot>
        </Table>
      </TableContainer>
    </Stack>
  );
};
