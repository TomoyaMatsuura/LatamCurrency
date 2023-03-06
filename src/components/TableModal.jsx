import {
  Table,
  Thead,
  Tbody,
  Tr,
  Td,
  Th,
  TableContainer
} from "@chakra-ui/react";

export const TableModal = (props) => {
  const { FinData } = props;
  console.log(FinData);
  return (
    <TableContainer color="gray.500">
      <Table>
        <Thead>
          <Tr>
            {FinData.Header.map((header) => (
              <Th key={header}>{header}</Th>
            ))}
          </Tr>
        </Thead>
        <Tbody>
          {FinData.Body.map((body) => (
            <Tr key={body}>
              {body.map((f) => (
                <Td key={f}>{f}</Td>
              ))}
            </Tr>
          ))}
        </Tbody>
      </Table>
    </TableContainer>
  );
};
