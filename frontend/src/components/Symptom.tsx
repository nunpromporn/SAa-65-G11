import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
// import { createStyles, makeStyles, Theme } from "@mui/material/styles";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Box from "@mui/material/Box";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import { SymptomInterface } from "../models/ISymptom";
import { format } from 'date-fns'

// const useStyles = makeStyles((theme: Theme) =>
//   createStyles({
//     container: {
//       marginTop: theme.spacing(2),
//     },
//     table: {
//       minWidth: 650,
//     },
//     tableSpace: {
//       marginTop: 20,
//     },
//   })
// );

function Symptoms() {
  // const classes = useStyles();
  const [symptoms, setSymptoms] = useState<SymptomInterface[]>([]);
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const getSymptoms = async () => {
    fetch(`${apiUrl}/symptoms`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
            setSymptoms(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getSymptoms();
  }, []);

  return (
    <div>
      <Container sx={{marginTop: 2}} maxWidth="md">
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              การบันทึกอาการ
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/symptom/create"
              variant="contained"
              color="primary"
            >
              บันทึกอาการ
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} sx={{ minWidth: 650}}>
          <Table sx={{marginTop: 2}} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="2%">
                  ลำดับ
                </TableCell>
                <TableCell align="center" width="18%">
                  ชื่อผู้ตรวจ
                </TableCell>
                <TableCell align="center" width="18%">
                  ชื่อผู้ป่วย
                </TableCell>
                <TableCell align="center" width="18%">
                  บริเวณที่เจ็บป่วย
                </TableCell>
                <TableCell align="center" width="15%">
                  แผนกที่ตรวจ
                </TableCell>
                <TableCell align="center" width="20%">
                  คำอธิบาย
                </TableCell>
                <TableCell align="center" width="30%">
                  วันที่และเวลา
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {symptoms.map((item: SymptomInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.User.Name}</TableCell>
                  <TableCell align="center">{item.Patient.Name}</TableCell>
                  <TableCell align="center">{item.Tenderness.Name}</TableCell>
                  <TableCell align="center">{item.Department.Name}</TableCell>
                  <TableCell align="center">{item.Explain}</TableCell>
                  <TableCell align="center">{format((new Date(item.SymptomTime)), 'dd MMMM yyyy hh:mm')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default Symptoms;