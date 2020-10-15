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
import { EntBookBorrow} from '../../api/models/EntBookBorrow';
const useStyles = makeStyles({
    table: {
        minWidth: 650,
    },
});

export default function ComponentsTable() {
    const classes = useStyles();
    const api = new DefaultApi();
    
    const [bookborrows, setBookBorrows] = useState<EntBookBorrow[]>(Array);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const getUsers = async () => {
            const res = await api.listBookborrow({ limit: 20, offset: 0 });
            setLoading(false);
            setBookBorrows(res);
        };
        getUsers();
    }, [loading]);
    
    return (
        <TableContainer component={Paper}>
            <Table className={classes.table} aria-label="simple table">
                <TableHead>
                    <TableRow>
                        <TableCell align="center">หมายเลขการยืม</TableCell>
                        <TableCell align="center">ผู้ยืม</TableCell>
                        <TableCell align="center">หนังสือ</TableCell>
                        <TableCell align="center">วัตถุประสงค์</TableCell>
                        <TableCell align="center">วันเวลาที่ยืม</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {bookborrows.map((item:any)  => (
                        <TableRow key={item.id}>
                            <TableCell align="center">{item.id}</TableCell>
                            <TableCell align="center">{item.edges.user.uSERNAME}</TableCell>
                            <TableCell align="center">{item.edges.book.bOOKNAME}</TableCell>
                            <TableCell align="center">{item.edges.purpose.pURPOSENAME}</TableCell>
                            <TableCell align="center">{item.bOOKINGDATE}</TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </TableContainer>
    );
}