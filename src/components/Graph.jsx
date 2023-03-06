import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
} from "chart.js";
import { Line } from "react-chartjs-2";

import { Box, Image, Stack, useDisclosure } from "@chakra-ui/react";
import { ModalAffiliates } from "./ModalAffiliates";

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
);

export const Graph = (props) => {
  const { currency, color, rateThisYear, rateLastYear, flag, FinData } = props;
  const options = {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      title: {
        display: true,
        text: currency,
        font: {
          size: 20
        }
      }
    }
  };

  const labels = [
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

  const data = {
    labels,
    datasets: [
      {
        label: "FY110",
        data: rateLastYear,
        borderColor: "rgba(0, 0, 0, 0.1)",
        backgroundColor: "rgba(0, 0, 0, 0.1)",
        borderWidth: 1
      },
      {
        label: "FY111",
        data: rateThisYear,
        borderColor: color,
        backgroundColor: "white",
        borderWidth: 3
      }
    ]
  };
  const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <>
      <Box
        w="400px"
        h="330px"
        bg="white"
        borderRadius="10px"
        shadow="md"
        p={2}
        m={2}
      >
        <Stack h="240px" spacing="-2">
          <Image
            src={flag}
            boxSize="100px"
            alt="国旗"
            m="auto"
            onClick={onOpen}
            _hover={{ cursor: "pointer", opacity: 0.7 }}
          />
          <Line options={options} data={data} />
        </Stack>
      </Box>
      <ModalAffiliates isOpen={isOpen} onClose={onClose} FinData={FinData} />
    </>
  );
};
