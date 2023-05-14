import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomePageComponent } from './home-page/home-page.component';
import { TuzilastvoComponent } from './Tuzilastvo-Komponente/tuzilastvo/tuzilastvo.component';
import { MupPageComponent } from './mup-page/mup-page.component';
import { SudstvoPageComponent } from './sudstvo_components/sudstvo-page/sudstvo-page.component';
import { LogoutComponent } from './logout/logout.component';


const routes: Routes = [

  {path: '', component: HomePageComponent},
  {path: 'tuzilastvo', component: TuzilastvoComponent},
  {path: 'mup', component: MupPageComponent},

  {path: 'sudstvo', component: SudstvoPageComponent},

  {path: 'logout', component: LogoutComponent}


];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
