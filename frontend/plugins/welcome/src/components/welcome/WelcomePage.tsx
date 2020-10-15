import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';

import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';

const WelcomePage: FC<{}> = () => {

 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`ระบบยืมหนังสือ`}
     ></Header>
     <Content>
       <ContentHeader title="รายการยืมหนังสือ">
         <Link component={RouterLink} to="/BookBorrow">
           <Button variant="contained" color="primary">
             Add New Borrow
           </Button>

           
         </Link>
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};

export default WelcomePage;