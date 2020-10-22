import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';

const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [mealplans, setMealPlans] = useState(Array);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const getMealPlans = async () => {
     const res = await api.listMealplan({ limit: 10, offset: 0 });
     setLoading(false);
     setMealPlans(res);
     console.log(res);
   };
   getMealPlans();
 }, [loading]);
 
 const deleteMealPlans = async (id: number) => {
   const res = await api.deleteMealplan({ id: id });
   setLoading(true);
 };
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">ไอดี</TableCell>
           <TableCell align="center">วันที่</TableCell>
           <TableCell align="center">เจ้าของ</TableCell>
           <TableCell align="center">มื้ออาหาร</TableCell>
           <TableCell align="center">อาหาร</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {mealplans.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.date}</TableCell>
             <TableCell align="center">{item.edges?.owner?.name}</TableCell>
             <TableCell align="center">{item.edges?.meals?.name}</TableCell>
             <TableCell align="center">{item.edges?.foods?.name}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                   deleteMealPlans(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 ลบออก
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
