import {NgModule} from '@angular/core';
import {LoginComponent} from './login/login.component';
import {DashboardComponent} from './dashboard/dashboard.component';
import {RouterModule, Routes} from '@angular/router';
import {AuthGuardGuard} from './services/guard/auth-guard.guard';
import {UserResolverService} from './services/user-resolver/user-resolver.service';

const routes: Routes = [
  {path: 'login', component: LoginComponent},
  {path: '', component: DashboardComponent, canActivate: [AuthGuardGuard], resolve: { user: UserResolverService}}
];

@NgModule({
  declarations: [],
  imports: [
    RouterModule.forRoot(routes)
  ],
  exports: [
    RouterModule
  ]
})
export class AppRoutingModule {
}
