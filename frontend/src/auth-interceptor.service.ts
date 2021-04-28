import { Injectable, Injector } from '@angular/core';
import { HttpInterceptor } from '@angular/common/http';
import {UserService} from './app/services/user/user.service';

import { ToastrService } from 'ngx-toastr';
import {Router} from '@angular/router';
import {throwError} from 'rxjs';


@Injectable()
export class AuthInterceptorService implements HttpInterceptor {

  constructor(private router: Router, private injector: Injector, private toastr: ToastrService) { }

  intercept(req, next): any {
    const userService = this.injector.get(UserService);

    if (userService.isTokenExpired()) {
      userService.logout();
      this.toastr.warning('Session timed out! Please login');
      this.router.navigateByUrl('/login');
      return throwError('Session Timed Out');
    }

    const userRequest = req.clone({
      // tslint:disable-next-line:max-line-length
      headers: req.headers.set('Authorization', 'Bearer ' + userService.getToken())
    });

    return next.handle(userRequest);
  }
}
