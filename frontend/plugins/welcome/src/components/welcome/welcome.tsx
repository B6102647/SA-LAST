import React, { useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import AccountBoxIcon from '@material-ui/icons/AccountBox';
import Autocomplete from '@material-ui/lab/Autocomplete';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
 
const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
   },
   margin: {
     margin: theme.spacing(3),
   },
   withoutLabel: {
     marginTop: theme.spacing(3),
   },
   textField: {
     width: '25ch',
   },
 }),
);
 
const initialUserState = {
 name: 'System Analysis and Design',
 age: 20,
};
 
export default function Create() {
 const classes = useStyles();
 const profile = { givenName: 'to Software Analysis 63' };
 const api = new DefaultApi();
 
 const [user, setUser] = useState(initialUserState);
 const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(true);
 
 const handleInputChange = (event: any) => {
   const { id, value } = event.target;
   setUser({ ...user, [id]: value });
 };
 
 const CreateUser = async () => {
   const res = await api.createUser({ user });
   setStatus(true);
   if (res.id != ''){
     setAlert(true);
   } else {
     setAlert(false);
   }
   const timer = setTimeout(() => {
     setStatus(false);
   }, 1000);
 };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`ระบบยืมหนังสือ`}
     ></Header>
     <Content>
       <ContentHeader title="">
         {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 This is a success alert — check it out!
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 This is a warning alert — check it out!
               </Alert>
             )}
           </div>
         ) : null}
       </ContentHeader>
       <div className={classes.root}>
         <form noValidate autoComplete="off">
         <tr align='center'>
                  <td></td><td></td><td align='center'>
                    <AccountBoxIcon style={{ fontSize: 200}}>AccountBoxIcon</AccountBoxIcon>
                    <br></br>
                    jaehyun@gmail.com
                    <br></br>
                    <u><font color="blue">logout </font> </u>
                  </td>
            </tr>
          <tr>
              <tr>
                <td width='170' align="left"> <font size="4">หนังสือ&nbsp;&nbsp;</font> 
                <FormControl
                  //fullWidth
                  className={classes.margin}
                  variant="outlined">
                    <Autocomplete
                      id="book"
                      options={bookName}
                      getOptionLabel={(option) => option.title}
                      style={{ width: 300}}
                      renderInput={(params) => <TextField {...params} label="Book" variant="outlined" />}
                    />
                </FormControl></td>
              </tr>
              <tr><h1></h1></tr>
              <tr>
                  <td width='170' align="left"> <font size="4">Email&nbsp;&nbsp;</font> 
                  <FormControl
                    //fullWidth
                    className={classes.margin}
                    variant="e-mail">
                      <Autocomplete
                        id="auto-complete"
                        options={datetime}
                        getOptionLabel={(option) => option.title}
                        style={{ width: 300}}
                        renderInput={(params) => <TextField {...params} label="jaehyun@gmail.com" variant="outlined" />}
                      />
                  </FormControl></td>
                </tr>
                <tr><h1></h1></tr>
                <tr>
                  <td width='170' align="left"> <font size="4">บรรณารักษ์ผู้ดูแล&nbsp;&nbsp;</font> 
                  <FormControl
                    //fullWidth
                    className={classes.margin}
                    variant="outlined">
                      <Autocomplete
                        id="librarian"
                        options={numberofday}
                        getOptionLabel={(option) => option.title}
                        style={{ width: 300}}
                        renderInput={(params) => <TextField {...params} label="Librarian" variant="outlined" />}
                      />
                  </FormControl></td>
                </tr>
                <tr><h1></h1></tr>
                <tr>
                  <td width='170' align="left"> <font size="4">เวลา&nbsp;&nbsp;</font>
                  <FormControl
                    //fullWidth
                    className={classes.margin}
                    variant="outlined">
                      <Autocomplete
                        id="date time"
                        options={datetime}
                        getOptionLabel={(option) => option.title}
                        style={{ width: 300}}
                        renderInput={(params) => <TextField {...params} label="Date Time" variant="outlined" />}
                      />
                  </FormControl></td>
                </tr>
                <br></br>
                <br></br>
                <tr>
                  <td width='170'> <font size="4"></font> 
                  <TextField
                      id="datetime-local"
                      label="Select DateTime"
                      type="datetime-local"
                      defaultValue="2017-05-24T10:30"
                      className={classes.textField}
                      InputLabelProps={{
                        shrink: true,
                      }}
                    /></td>
                </tr>
                
            </tr> 
            
           <div className={classes.margin}>
             <Button
               onClick={() => {
                 CreateUser();
               }}
               variant="contained"
               color="primary"
             >
               Submit
             </Button>
             <Button
               style={{ marginLeft: 20 }}
               component={RouterLink}
               to="/"
               variant="contained"
             >
               Back
             </Button>
           </div>
         </form>
       </div>
     </Content>
   </Page>
 );
}

const mail = [
  { title: 'jaehyun@gmail.com'},
];
const datetime = [
  { title: '2020-08-28 10:00PM'},
];

const bookName = [
  { title: 'Microprocessor for Beginner'},
  { title: 'System Analysis'},
  { title: 'Solo Leveling Novel'},
];
const numberofday = [
  { title: 'Sompong Deejai'},
  { title: 'Meetang Makmak'},
  { title: 'Pensri Dangdam'},
];