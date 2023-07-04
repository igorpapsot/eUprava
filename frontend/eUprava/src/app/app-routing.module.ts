import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomePageComponent } from './home-page/home-page.component';
import { TuzilastvoComponent } from './Tuzilastvo-Komponente/tuzilastvo/tuzilastvo.component';
import { MupPageComponent } from './mup-page/mup-page.component';
import { SudstvoPageComponent } from './sudstvo_components/sudstvo-page/sudstvo-page.component';
import { LogoutComponent } from './logout/logout.component';
import { CreatePoternicaComponent } from './sudstvo_components/poternice/create-poternica/create-poternica.component';
import { CreateRocisteComponent } from './sudstvo_components/rocista/create-rociste/create-rociste.component';
import { ViewPoterniceComponent } from './sudstvo_components/poternice/view-poternice/view-poternice.component';
import { ViewRocisteComponent } from './sudstvo_components/rocista/view-rociste/view-rociste.component';
import { CreateSudijaComponent } from './sudstvo_components/sudija/create-sudija/create-sudija.component';
import { ViewSudijeComponent } from './sudstvo_components/sudija/view-sudije/view-sudije.component';
import { ViewOptuzniceComponent } from './sudstvo_components/optuznice/view-optuznice/view-optuznice.component';


const routes: Routes = [

  {path: '', component: HomePageComponent},
  {path: 'tuzilastvo', component: TuzilastvoComponent},
  {path: 'mup', component: MupPageComponent},

  {path: 'sudstvo', component: SudstvoPageComponent},
  {path: 'create-poternica', component: CreatePoternicaComponent},
  {path: 'create-rociste', component: CreateRocisteComponent},
  {path: 'view-poternice', component: ViewPoterniceComponent},
  {path: 'view-rociste', component: ViewRocisteComponent},
  {path: 'register-sudija', component: CreateSudijaComponent},
  {path: 'view-sudije', component: ViewSudijeComponent},
  {path: 'view-optuznice', component: ViewOptuzniceComponent},
  {path: 'logout', component: LogoutComponent}


];

@NgModule({
  imports: [RouterModule.forRoot(routes, {onSameUrlNavigation:"reload"})],
  exports: [RouterModule]
})
export class AppRoutingModule { }
