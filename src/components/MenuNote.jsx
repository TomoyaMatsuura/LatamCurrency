import {
  Menu,
  MenuButton,
  ChevronDownIcon,
  MenuList,
  MenuItem
} from "@chakra-ui/react";

export const MenuNote = () => {
  return (
    <Menu>
      <MenuButton paddingRight="7" rightIcon={<ChevronDownIcon />}>
        使用方法
      </MenuButton>
      <MenuList>
        <MenuItem bgColor="white" color="gray.500">
          Excelはxlsx形式
        </MenuItem>
        <MenuItem bgColor="white" color="gray.500">
          FYの部分をクリックすると表示非表示が可能
        </MenuItem>
        <MenuItem bgColor="white" color="gray.500">
          国旗をクリックすると各国経済指標が表示
        </MenuItem>
      </MenuList>
    </Menu>
  );
};
