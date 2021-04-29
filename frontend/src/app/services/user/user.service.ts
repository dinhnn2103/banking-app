import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {BehaviorSubject, Observable, of} from 'rxjs';
import {Router} from '@angular/router';
import {LoginInfo} from '../../login/loginInfo';
import jwtDecode from 'jwt-decode';
import {ToastrService} from 'ngx-toastr';

const httpOptions = {
  headers: new HttpHeaders({
    'Access-Control-Allow-Origin': '*',
    'Content-Type': 'application/json'
  })
};

export const TOKEN_NAME = 'jwt_token';
export const USER_ID_KEY = 'userId';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  url: any = 'http://localhost:4200/api/';
  errorSubject: any = new BehaviorSubject<any>(null);
  errorMessage: any = this.errorSubject.asObservable();

  userSubject: any = new BehaviorSubject<any>(null);
  user: any = this.userSubject.asObservable();

  private loggedIn = new BehaviorSubject<boolean>(false);

  get isLoggedIn(): Observable<boolean> {
    return this.loggedIn.asObservable();
  }

  constructor(
    private http: HttpClient,
    private router: Router,
    private toastr: ToastrService
  ) {
  }

  getToken(): string {
    return localStorage.getItem(TOKEN_NAME);
  }

  setToken(token: string): void {
    localStorage.setItem(TOKEN_NAME, token);
  }

  getTokenExpirationDate(token: string): Date {

    const decoded: any = jwtDecode(token);

    if (decoded.expiry === undefined) {
      return null;
    }

    const date = new Date(0);
    date.setUTCSeconds(decoded.expiry);

    return date;
  }

  isTokenExpired(token?: string): boolean {

    if (!token) {
      token = this.getToken();
    }
    if (!token) {
      return null;
    }

    const date = this.getTokenExpirationDate(token);
    if (date === undefined) {
      return false;
    }

    return !(date.valueOf() > new Date().valueOf());
  }

  login(loginInfo: LoginInfo): any {

    return this.http.post(`${this.url}login`, loginInfo, httpOptions).toPromise().then((res: any) => {

      if (res && res.jwt) {
        localStorage.setItem(TOKEN_NAME, res.jwt);
        this.errorSubject.next(null);
        if (res.data) {
          this.userSubject.next(res.data);
          localStorage.setItem(USER_ID_KEY, res.data.ID);
        }
        this.loggedIn.next(true);
        this.router.navigateByUrl('');
      } else if (res.message) {
        this.toastr.error('Authentication failed (Username / password invalid)');
        console.log(res.message);
        // this.errorSubject.next(res.message);
      }

    }).catch(err => {
      this.toastr.error('Can\'t login, please check backend');
      console.log(err);
    });
  }

  getUser(): any {
    const userId = localStorage.getItem(USER_ID_KEY);
    return this.http.get(`${this.url}user/${userId}`);
  }

  logout(): void {
    this.loggedIn.next(false);
    localStorage.removeItem(TOKEN_NAME);
    this.router.navigateByUrl('login');
  }

  updateLoggedIn(): void {
    if (!localStorage.getItem(TOKEN_NAME) || this.isTokenExpired()) {
      this.loggedIn.next(false);
    } else {
      this.loggedIn.next(true);
    }
  }
}
