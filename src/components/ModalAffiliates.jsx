import {
  Modal,
  ModalContent,
  ModalOverlay,
  ModalHeader,
  ModalCloseButton
} from "@chakra-ui/react";

import { TableModal } from "./TableModal";

export const ModalAffiliates = (props) => {
  const { isOpen, onClose, FinData } = props;
  return (
    <Modal isCentered="true" size={"3xl"} isOpen={isOpen} onClose={onClose}>
      <ModalOverlay />
      <ModalContent textAlign="center">
        <ModalHeader color="gray.500">{FinData.Tag}</ModalHeader>
        <TableModal FinData={FinData} />
        <ModalCloseButton />
      </ModalContent>
    </Modal>
  );
};
