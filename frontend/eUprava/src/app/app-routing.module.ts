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
import { PoterniceListComponent } from './sudstvo_components/poternice/poternice-list/poternice-list.component';
import { RocisteListComponent } from './sudstvo_components/rocista/rociste-list/rociste-list.component';


const routes: Routes = [

  {path: '', component: HomePageComponent},
  {path: 'tuzilastvo', component: TuzilastvoComponent},
  {path: 'mup', component: MupPageComponent},

  {path: 'sudstvo', component: SudstvoPageComponent},
  {path: 'create-poternica', component: CreatePoternicaComponent},
  {path: 'create-rociste', component: CreateRocisteComponent},
  {path: 'view-poternice', component: ViewPoterniceComponent},
  {path: 'view-rociste', component: ViewRocisteComponent},
  {path: 'poternice-list/:id', component: PoterniceListComponent},
  {path: 'rociste-list/:id', component: RocisteListComponent},
  {path: 'logout', component: LogoutComponent}


];

@NgModule({
  imports: [RouterModule.forRoot(routes, {onSameUrlNavigation:"reload"})],
  exports: [RouterModule]
})
export class AppRoutingModule { }
