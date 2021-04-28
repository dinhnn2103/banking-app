import { Injectable } from '@angular/core';
import {CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot, UrlTree, Router} from '@angular/router';
import {TOKEN_NAME, USER_ID_KEY, UserService} from '../user/user.service';
import {Observable, of} from 'rxjs';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {catchError, map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AuthGuardGuard implements CanActivate {

  url = 'http://localhost:4200/api/';

  constructor(
    public http: HttpClient,
    public router: Router,
    public userService: UserService
  ) {}

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): Observable<boolean> {

    const userId = localStorage.getItem(USER_ID_KEY);
    const jwtToken = localStorage.getItem(TOKEN_NAME);

    this.userService.updateLoggedIn();

    if (userId && jwtToken) {
      // authorization succeed
      return of(true);
    } else {
      this.router.navigateByUrl('login');
      return of(false);
    }

  }
}
