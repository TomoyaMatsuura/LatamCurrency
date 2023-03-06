import { Button } from "@chakra-ui/react";
import ExcelJS from "exceljs";

import { MBR } from "./affiliates/MBR";
import { MMX } from "./affiliates/MMX";
import { MCL } from "./affiliates/MCL";
import { MAR } from "./affiliates/MAR";
import { MPE } from "./affiliates/MPE";
import { MCO } from "./affiliates/MCO";

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

async function makeExcel() {
  console.log(affiliates);
  //ワークブックを作成する
  let workbook = new ExcelJS.Workbook();
  //ワークシートを作成する
  let worksheet = workbook.addWorksheet("中南米為替", {});
  //1列目の幅を調整
  worksheet.getColumn(1).width = 7;
  //列を取得し、通貨を記入
  let sheet_row = worksheet.getRow(1);
  for (let i = 1; i < 13; i++) {
    sheet_row.getCell(i + 1).value = month[i - 1];
  }
  sheet_row = worksheet.getRow(2);
  sheet_row.getCell(1).value = affiliates[0].currency;
  for (let ii = 0; ii < 13; ii++) {
    sheet_row.getCell(ii + 2).value = affiliates[0].rateThisYear[ii];
  }
  sheet_row = worksheet.getRow(3);
  sheet_row.getCell(1).value = affiliates[1].currency;
  for (let ii = 0; ii < 13; ii++) {
    sheet_row.getCell(ii + 2).value = affiliates[1].rateThisYear[ii];
  }
  sheet_row = worksheet.getRow(4);
  sheet_row.getCell(1).value = affiliates[2].currency;
  for (let ii = 0; ii < 13; ii++) {
    sheet_row.getCell(ii + 2).value = affiliates[2].rateThisYear[ii];
  }
  sheet_row = worksheet.getRow(5);
  sheet_row.getCell(1).value = affiliates[3].currency;
  for (let ii = 0; ii < 13; ii++) {
    sheet_row.getCell(ii + 2).value = affiliates[3].rateThisYear[ii];
  }
  sheet_row = worksheet.getRow(6);
  sheet_row.getCell(1).value = affiliates[4].currency;
  for (let ii = 0; ii < 13; ii++) {
    sheet_row.getCell(ii + 2).value = affiliates[4].rateThisYear[ii];
  }
  sheet_row = worksheet.getRow(7);
  sheet_row.getCell(1).value = affiliates[5].currency;
  for (let ii = 0; ii < 13; ii++) {
    sheet_row.getCell(ii + 2).value = affiliates[5].rateThisYear[ii];
  }

  //エクセルファイルを生成する
  let uint8Array = await workbook.xlsx.writeBuffer(); //xlsxの場合
  let blob = new Blob([uint8Array], { type: "application/octet-binary" });

  //Excelファイルダウンロード
  let link = document.createElement("a");
  link.href = window.URL.createObjectURL(blob);
  link.download = "中南米為替.xlsx";
  link.click();
}

export const PutExcelButton = (props) => {
  return (
    <Button
      w="100px"
      bgColor="teal.500"
      color="white"
      variant="contained"
      onClick={makeExcel}
    >
      Excel出力
    </Button>
  );
};
