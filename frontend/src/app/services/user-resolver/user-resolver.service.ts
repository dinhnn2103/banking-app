import {ActivatedRouteSnapshot, Resolve, Router, RouterStateSnapshot} from '@angular/router';
import {UserService} from '../user/user.service';
import {Injectable} from '@angular/core';
import {EMPTY, Observable} from 'rxjs';
import {catchError} from 'rxjs/operators';
import {ToastrService} from 'ngx-toastr';

@Injectable({
  providedIn: 'root'
})
export class UserResolverService implements Resolve<any> {

  constructor(private user: UserService, private toastr: ToastrService) {
  }

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): any | Observable<never> {
    return this.user.getUser().pipe(catchError(err => {
      this.user.logout();
      this.toastr.error('Can\'t get user info');
      console.log(err);
      return EMPTY;
    }));
  }

}
