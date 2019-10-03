import { NgModule, ModuleWithProviders } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { MovieListComponent } from './components/movie-list/movie-list.component';

/** Routes */
const routes: Routes = [
  {
    path: '',
    redirectTo: 'movie-list',
    pathMatch: 'full'
  },
  {
    path: 'movie-list',
    component: MovieListComponent
  }
];


@NgModule({
  imports: [RouterModule.forRoot(routes,{ useHash: true, onSameUrlNavigation: 'reload'})],
  exports: [RouterModule]
})
export class AppRoutingModule { }
