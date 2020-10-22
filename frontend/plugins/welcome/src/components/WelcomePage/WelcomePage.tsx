import React, { FC } from 'react';
import { Grid } from '@material-ui/core';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import ComponentsTable from '../Table';

const HeaderCustom = {
  minHeight: '100px',
};

const WelcomePage: FC<{}> = () => {

  return (
    <Page theme={pageTheme.home}>
      <Header
        style={HeaderCustom}
        title={'MEAL PLANNER'}
      ></Header>
      <Content>
        <ContentHeader title="Meal Plan">
        </ContentHeader>
        <Grid container>
          <Grid item xs={12}>
          <ComponentsTable></ComponentsTable>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
};

export default WelcomePage;
