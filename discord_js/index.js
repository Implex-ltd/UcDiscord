"use strict";
(this.webpackChunkdiscord_app = this.webpackChunkdiscord_app || []).push([
    [98847], {
        101566: (e, t, n) => {
            n.d(t, {
                Z: () => J,
                $: () => V
            });
            var r = n(495414),
                o = n(409479),
                i = n(95508),
                l = n(598143),
                a = n(257695),
                u = n(595882),
                c = n(717187),
                s = n.n(c),
                f = n(496486),
                d = n.n(f),
                p = n(296602),
                m = n(493254);

            function y(e, t) {
                (null == t || t > e.length) && (t = e.length);
                for (var n = 0, r = new Array(t); n < t; n++) r[n] = e[n];
                return r
            }

            function h(e, t) {
                if (!(e instanceof t)) throw new TypeError("Cannot call a class as a function")
            }

            function v(e, t, n) {
                t in e ? Object.defineProperty(e, t, {
                    value: n,
                    enumerable: !0,
                    configurable: !0,
                    writable: !0
                }) : e[t] = n;
                return e
            }

            function b(e) {
                b = Object.setPrototypeOf ? Object.getPrototypeOf : function (e) {
                    return e.__proto__ || Object.getPrototypeOf(e)
                };
                return b(e)
            }

            function g(e) {
                for (var t = 1; t < arguments.length; t++) {
                    var n = null != arguments[t] ? arguments[t] : {},
                        r = Object.keys(n);
                    "function" == typeof Object.getOwnPropertySymbols && (r = r.concat(Object.getOwnPropertySymbols(n).filter((function (e) {
                        return Object.getOwnPropertyDescriptor(n, e).enumerable
                    }))));
                    r.forEach((function (t) {
                        v(e, t, n[t])
                    }))
                }
                return e
            }

            function _(e, t) {
                t = null != t ? t : {};
                Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : function (e, t) {
                    var n = Object.keys(e);
                    if (Object.getOwnPropertySymbols) {
                        var r = Object.getOwnPropertySymbols(e);
                        t && (r = r.filter((function (t) {
                            return Object.getOwnPropertyDescriptor(e, t).enumerable
                        })));
                        n.push.apply(n, r)
                    }
                    return n
                }(Object(t)).forEach((function (n) {
                    Object.defineProperty(e, n, Object.getOwnPropertyDescriptor(t, n))
                }));
                return e
            }

            function O(e, t) {
                return !t || "object" !== E(t) && "function" != typeof t ? function (e) {
                    if (void 0 === e) throw new ReferenceError("this hasn't been initialised - super() hasn't been called");
                    return e
                }(e) : t
            }

            function w(e, t) {
                w = Object.setPrototypeOf || function (e, t) {
                    e.__proto__ = t;
                    return e
                };
                return w(e, t)
            }

            function S(e) {
                return function (e) {
                    if (Array.isArray(e)) return y(e)
                }(e) || function (e) {
                    if ("undefined" != typeof Symbol && null != e[Symbol.iterator] || null != e["@@iterator"]) return Array.from(e)
                }(e) || function (e, t) {
                    if (!e) return;
                    if ("string" == typeof e) return y(e, t);
                    var n = Object.prototype.toString.call(e).slice(8, -1);
                    "Object" === n && e.constructor && (n = e.constructor.name);
                    if ("Map" === n || "Set" === n) return Array.from(n);
                    if ("Arguments" === n || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n)) return y(e, t)
                }(e) || function () {
                    throw new TypeError("Invalid attempt to spread non-iterable instance.\\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")
                }()
            }
            var E = function (e) {
                return e && "undefined" != typeof Symbol && e.constructor === Symbol ? "symbol" : typeof e
            };

            function I(e) {
                var t = function () {
                    if ("undefined" == typeof Reflect || !Reflect.construct) return !1;
                    if (Reflect.construct.sham) return !1;
                    if ("function" == typeof Proxy) return !0;
                    try {
                        Boolean.prototype.valueOf.call(Reflect.construct(Boolean, [], (function () {})));
                        return !0
                    } catch (e) {
                        return !1
                    }
                }();
                return function () {
                    var n, r = b(e);
                    if (t) {
                        var o = b(this).constructor;
                        n = Reflect.construct(r, arguments, o)
                    } else n = r.apply(this, arguments);
                    return O(this, n)
                }
            }
            var x = new p.Z("UploaderBase.tsx"),
                A = function (e) {
                    ! function (e, t) {
                        if ("function" != typeof t && null !== t) throw new TypeError("Super expression must either be null or a function");
                        e.prototype = Object.create(t && t.prototype, {
                            constructor: {
                                value: e,
                                writable: !0,
                                configurable: !0
                            }
                        });
                        t && w(e, t)
                    }(n, e);
                    var t = I(n);

                    function n(e) {
                        var r, o, i = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : "POST",
                            l = arguments.length > 2 ? arguments[2] : void 0;
                        h(this, n);
                        (r = t.call(this))._token = "";
                        r._lastUpdate = 0;
                        r._loaded = 0;
                        r._aborted = !1;
                        r._errored = !1;
                        r._raiseEndpointErrors = !1;
                        r.alreadyStarted = !1;
                        r._handleStart = function (e) {
                            r._cancel = e;
                            r.alreadyStarted || r.emit("start", r._file);
                            r.alreadyStarted = !0
                        };
                        r._handleProgress = function (e, t, n) {
                            var o = Date.now(),
                                i = (0, u.S)(e, t),
                                l = Math.floor((e - r._loaded) / ((o - r._lastUpdate) / 1e3));
                            if (null != n) {
                                var a;
                                null === (a = r._file.items) || void 0 === a || a.forEach((function (e) {
                                    e.item.progress = n[e.id]
                                }))
                            }
                            r._lastUpdate = o;
                            r._loaded = e;
                            r._file = _(g({}, r._file), {
                                currentSize: t,
                                progress: i,
                                rate: l
                            });
                            r.emit("progress", r._file)
                        };
                        r._handleException = function (e) {
                            r._handleError({
                                reason: {
                                    type: m.xi.ERROR_SOURCE_UNKNOWN,
                                    msg: e.toString()
                                }
                            })
                        };
                        r._handleAborted = function () {
                            r.clearProcessingMessageInterval()
                        };
                        r._handleError = function (e) {
                            var t = e.code,
                                n = e.reason,
                                o = e.body;
                            r.clearProcessingMessageInterval();
                            if (!r._aborted) {
                                r._errored = !0;
                                x.log("_handleError: ".concat(t, " (").concat(JSON.stringify(n), ") for ").concat(r.id));
                                r.emit("error", r._file, t, o, n);
                                r.removeAllListeners()
                            }
                        };
                        r._handleComplete = function (e) {
                            r.clearProcessingMessageInterval();
                            x.log("_handleComplete for ".concat(r.id));
                            r.emit("complete", r._file, e);
                            r.removeAllListeners()
                        };
                        r.id = d().uniqueId("Uploader");
                        r._url = e;
                        r._method = i;
                        r._raiseEndpointErrors = null !== (o = null == l ? void 0 : l.raiseEndpointErrors) && void 0 !== o && o;
                        return r
                    }
                    var r = n.prototype;
                    r._addAttachmentsToPayload = function (e, t, n) {
                        var r = g({}, e),
                            o = S(d().get(r, t, [])).concat(S(n));
                        return d().set(r, t, o)
                    };
                    r.clearProcessingMessageInterval = function () {
                        if (null != this.processingMessageChangeInterval) {
                            clearInterval(this.processingMessageChangeInterval);
                            this.processingMessageChangeInterval = void 0
                        }
                    };
                    r.cancel = function () {
                        x.log("cancel() for ".concat(this.id));
                        this._aborted = !0;
                        null != this._cancel && this._cancel();
                        this._handleComplete()
                    };
                    r.cancelItem = function (e) {
                        throw new Error("cancelItem() is not implemented on UploaderBase; must implement cancelItem() on subclass")
                    };
                    r.upload = function (e, t, n) {
                        if (null != this._cancel) throw new Error("Uploader.upload(...): An upload is already in progress.");
                        this._lastUpdate = Date.now();
                        this._loaded = 0;
                        this._file = {
                            id: this.id,
                            name: e.name,
                            currentSize: 0,
                            totalPreCompressionSize: 0,
                            compressionProgress: 0,
                            progress: 0,
                            rate: 0,
                            hasImage: !1,
                            hasVideo: !1,
                            attachmentsCount: 0,
                            draftContent: null == t ? void 0 : t.content,
                            channelId: null == t ? void 0 : t.channel_id,
                            items: n
                        }
                    };
                    return n
                }(s()),
                T = n(2590),
                C = n(473708);

            function P(e, t) {
                (null == t || t > e.length) && (t = e.length);
                for (var n = 0, r = new Array(t); n < t; n++) r[n] = e[n];
                return r
            }

            function M(e, t, n, r, o, i, l) {
                try {
                    var a = e[i](l),
                        u = a.value
                } catch (e) {
                    n(e);
                    return
                }
                a.done ? t(u) : Promise.resolve(u).then(r, o)
            }

            function j(e) {
                return function () {
                    var t = this,
                        n = arguments;
                    return new Promise((function (r, o) {
                        var i = e.apply(t, n);

                        function l(e) {
                            M(i, r, o, l, a, "next", e)
                        }

                        function a(e) {
                            M(i, r, o, l, a, "throw", e)
                        }
                        l(void 0)
                    }))
                }
            }

            function N(e, t) {
                if (!(e instanceof t)) throw new TypeError("Cannot call a class as a function")
            }

            function U(e, t, n) {
                t in e ? Object.defineProperty(e, t, {
                    value: n,
                    enumerable: !0,
                    configurable: !0,
                    writable: !0
                }) : e[t] = n;
                return e
            }

            function R(e) {
                R = Object.setPrototypeOf ? Object.getPrototypeOf : function (e) {
                    return e.__proto__ || Object.getPrototypeOf(e)
                };
                return R(e)
            }

            function k(e) {
                for (var t = 1; t < arguments.length; t++) {
                    var n = null != arguments[t] ? arguments[t] : {},
                        r = Object.keys(n);
                    "function" == typeof Object.getOwnPropertySymbols && (r = r.concat(Object.getOwnPropertySymbols(n).filter((function (e) {
                        return Object.getOwnPropertyDescriptor(n, e).enumerable
                    }))));
                    r.forEach((function (t) {
                        U(e, t, n[t])
                    }))
                }
                return e
            }

            function L(e, t) {
                t = null != t ? t : {};
                Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : function (e, t) {
                    var n = Object.keys(e);
                    if (Object.getOwnPropertySymbols) {
                        var r = Object.getOwnPropertySymbols(e);
                        t && (r = r.filter((function (t) {
                            return Object.getOwnPropertyDescriptor(e, t).enumerable
                        })));
                        n.push.apply(n, r)
                    }
                    return n
                }(Object(t)).forEach((function (n) {
                    Object.defineProperty(e, n, Object.getOwnPropertyDescriptor(t, n))
                }));
                return e
            }

            function D(e, t) {
                return !t || "object" !== B(t) && "function" != typeof t ? function (e) {
                    if (void 0 === e) throw new ReferenceError("this hasn't been initialised - super() hasn't been called");
                    return e
                }(e) : t
            }

            function z(e, t) {
                z = Object.setPrototypeOf || function (e, t) {
                    e.__proto__ = t;
                    return e
                };
                return z(e, t)
            }

            function Z(e) {
                return function (e) {
                    if (Array.isArray(e)) return P(e)
                }(e) || function (e) {
                    if ("undefined" != typeof Symbol && null != e[Symbol.iterator] || null != e["@@iterator"]) return Array.from(e)
                }(e) || function (e, t) {
                    if (!e) return;
                    if ("string" == typeof e) return P(e, t);
                    var n = Object.prototype.toString.call(e).slice(8, -1);
                    "Object" === n && e.constructor && (n = e.constructor.name);
                    if ("Map" === n || "Set" === n) return Array.from(n);
                    if ("Arguments" === n || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n)) return P(e, t)
                }(e) || function () {
                    throw new TypeError("Invalid attempt to spread non-iterable instance.\\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")
                }()
            }
            var B = function (e) {
                return e && "undefined" != typeof Symbol && e.constructor === Symbol ? "symbol" : typeof e
            };

            function F(e) {
                var t = function () {
                    if ("undefined" == typeof Reflect || !Reflect.construct) return !1;
                    if (Reflect.construct.sham) return !1;
                    if ("function" == typeof Proxy) return !0;
                    try {
                        Boolean.prototype.valueOf.call(Reflect.construct(Boolean, [], (function () {})));
                        return !0
                    } catch (e) {
                        return !1
                    }
                }();
                return function () {
                    var n, r = R(e);
                    if (t) {
                        var o = R(this).constructor;
                        n = Reflect.construct(r, arguments, o)
                    } else n = r.apply(this, arguments);
                    return D(this, n)
                }
            }
            var G = function (e, t) {
                    var n, r, o, i, l = {
                        label: 0,
                        sent: function () {
                            if (1 & o[0]) throw o[1];
                            return o[1]
                        },
                        trys: [],
                        ops: []
                    };
                    return i = {
                        next: a(0),
                        throw: a(1),
                        return: a(2)
                    }, "function" == typeof Symbol && (i[Symbol.iterator] = function () {
                        return this
                    }), i;

                    function a(i) {
                        return function (a) {
                            return function (i) {
                                if (n) throw new TypeError("Generator is already executing.");
                                for (; l;) try {
                                    if (n = 1, r && (o = 2 & i[0] ? r.return : i[0] ? r.throw || ((o = r.return) && o.call(r), 0) : r.next) && !(o = o.call(r, i[1])).done) return o;
                                    (r = 0, o) && (i = [2 & i[0], o.value]);
                                    switch (i[0]) {
                                        case 0:
                                        case 1:
                                            o = i;
                                            break;
                                        case 4:
                                            l.label++;
                                            return {
                                                value: i[1], done: !1
                                            };
                                        case 5:
                                            l.label++;
                                            r = i[1];
                                            i = [0];
                                            continue;
                                        case 7:
                                            i = l.ops.pop();
                                            l.trys.pop();
                                            continue;
                                        default:
                                            if (!(o = l.trys, o = o.length > 0 && o[o.length - 1]) && (6 === i[0] || 2 === i[0])) {
                                                l = 0;
                                                continue
                                            }
                                            if (3 === i[0] && (!o || i[1] > o[0] && i[1] < o[3])) {
                                                l.label = i[1];
                                                break
                                            }
                                            if (6 === i[0] && l.label < o[1]) {
                                                l.label = o[1];
                                                o = i;
                                                break
                                            }
                                            if (o && l.label < o[2]) {
                                                l.label = o[2];
                                                l.ops.push(i);
                                                break
                                            }
                                            o[2] && l.ops.pop();
                                            l.trys.pop();
                                            continue
                                    }
                                    i = t.call(e, l)
                                } catch (e) {
                                    i = [6, e];
                                    r = 0
                                } finally {
                                    n = o = 0
                                }
                                if (5 & i[0]) throw i[1];
                                return {
                                    value: i[0] ? i[1] : void 0,
                                    done: !0
                                }
                            }([i, a])
                        }
                    }
                },
                H = new r.Z("CloudUploaderBase.tsx"),
                J = function (e) {
                    ! function (e, t) {
                        if ("function" != typeof t && null !== t) throw new TypeError("Super expression must either be null or a function");
                        e.prototype = Object.create(t && t.prototype, {
                            constructor: {
                                value: e,
                                writable: !0,
                                configurable: !0
                            }
                        });
                        t && z(e, t)
                    }(n, e);
                    var t = F(n);

                    function n() {
                        N(this, n);
                        var e;
                        (e = t.apply(this, arguments)).files = [];
                        return e
                    }
                    var r = n.prototype;
                    r._fileSize = function () {
                        return this.files.reduce((function (e, t) {
                            var n;
                            return e += null !== (n = t.currentSize) && void 0 !== n ? n : 0
                        }), 0)
                    };
                    r.compressAndCheckFileSize = function () {
                        var e = this;
                        return j((function () {
                            var t, n, r, o, i, u, c, s, f, d, p, y, h;
                            return G(this, (function (v) {
                                switch (v.label) {
                                    case 0:
                                        r = (0, a.F)(null === (t = e.files[0]) || void 0 === t || null === (n = t.item) || void 0 === n ? void 0 : n.target);
                                        if (e.files.length > r.getMaxAttachmentsCount()) {
                                            H.log("Too many attachments for ".concat(e.id));
                                            e._handleError({
                                                code: T.evJ.TOO_MANY_ATTACHMENTS
                                            });
                                            return [2, !1]
                                        }
                                        return [3, 10];
                                    case 1:
                                        v.trys.push([1, 8, 9, 10]);
                                        c = e.files[Symbol.iterator]();
                                        v.label = 2;
                                    case 2:
                                        if (o = (s = c.next()).done) return [3, 7];
                                        f = s.value;
                                        v.label = 3;
                                    case 3:
                                        v.trys.push([3, 5, , 6]);
                                        return [4, f.reactNativeCompressAndExtractData()];
                                    case 4:
                                        v.sent();
                                        if (f.status === l.m.CANCELED) {
                                            H.log("compressAndCheckFileSize() file has been cancelled for compression - ".concat(f.id));
                                            return [3, 6]
                                        }
                                        if (0 === (null !== (d = f.currentSize) && void 0 !== d ? d : 0)) {
                                            e._handleError({
                                                code: T.evJ.ENTITY_EMPTY
                                            });
                                            return [2, !1]
                                        }
                                        if ((null !== (p = f.currentSize) && void 0 !== p ? p : 0) > r.getMaxFileSize(f.channelId)) {
                                            e._handleError({
                                                code: T.evJ.ENTITY_TOO_LARGE,
                                                reason: {
                                                    type: m.xi.POSTCOMPRESSION_INDIVIDUAL_FILE_TOO_LARGE
                                                }
                                            });
                                            return [2, !1]
                                        }
                                        return [3, 6];
                                    case 5:
                                        y = v.sent();
                                        e._handleException(y);
                                        return [2, !1];
                                    case 6:
                                        o = !0;
                                        return [3, 2];
                                    case 7:
                                        return [3, 10];
                                    case 8:
                                        h = v.sent();
                                        i = !0;
                                        u = h;
                                        return [3, 10];
                                    case 9:
                                        try {
                                            o || null == c.return || c.return()
                                        } finally {
                                            if (i) throw u
                                        }
                                        return [7];
                                    case 10:
                                        if (e._fileSize() > r.getMaxTotalAttachmentSize()) {
                                            e._handleError({
                                                code: T.evJ.ENTITY_TOO_LARGE,
                                                reason: {
                                                    type: m.xi.POSTCOMPRESSION_SUM_TOO_LARGE
                                                }
                                            });
                                            return [2, !1]
                                        }
                                        return [2, !0]
                                }
                            }))
                        }))()
                    };
                    r._filesTooLarge = function () {
                        return this.files.some((function (e) {
                            return e.error === T.evJ.ENTITY_TOO_LARGE
                        }))
                    };
                    r.setUploadingTextForUI = function () {
                        var e = 1 === this.files.length && null != this.files[0].filename ? this.files[0].filename : C.Z.Messages.UPLOADING_FILES.format({
                                count: this.files.length
                            }),
                            t = this.files.some((function (e) {
                                return e.isImage
                            })),
                            n = this.files.some((function (e) {
                                return e.isVideo
                            })),
                            r = this._fileSize();
                        H.log("setUploadingTextForUI - total content: ".concat(r, " bytes and ").concat(this.files.length, " attachments for ").concat(this.id));
                        this._file = L(k({}, this._file), {
                            totalPostCompressionSize: r,
                            currentSize: r,
                            name: e,
                            hasVideo: n,
                            hasImage: t,
                            attachmentsCount: this.files.length,
                            items: this.files
                        })
                    };
                    r._recomputeProgress = function () {
                        var e, t = this._recomputeProgressTotal(),
                            n = t.loaded,
                            r = t.total;
                        (0, o.Dn)() && (e = this._recomputeProgressByFile());
                        this._handleProgress(n, r, e)
                    };
                    r._recomputeProgressTotal = function () {
                        var e = this._fileSize();
                        return {
                            loaded: this.files.reduce((function (e, t) {
                                var n;
                                return e += null !== (n = t.loaded) && void 0 !== n ? n : 0
                            }), 0),
                            total: e
                        }
                    };
                    r._recomputeProgressByFile = function () {
                        var e = {};
                        this.files.forEach((function (t) {
                            e[t.id] = (0, u.S)(t.loaded, t.currentSize)
                        }));
                        return e
                    };
                    r.cancel = function () {
                        H.log("Cancel called for ".concat(this.id));
                        if (!this._aborted) {
                            this._aborted = !0;
                            null != this._cancel && this._cancel();
                            this.files.forEach((function (e) {
                                return e.cancel()
                            }));
                            this._handleComplete()
                        }
                    };
                    r.cancelItem = function (e) {
                        var t = this;
                        return j((function () {
                            var n, r;
                            return G(this, (function (o) {
                                switch (o.label) {
                                    case 0:
                                        H.log("Cancel called for ".concat(t.id, " for item ").concat(e));
                                        if (null == (n = t.files.find((function (t) {
                                                return t.id === e
                                            })))) return [2];
                                        if (n.status === l.m.CANCELED) return [2];
                                        r = t.files.indexOf(n);
                                        t.files = Z(t.files.slice(0, r)).concat(Z(t.files.slice(r + 1)));
                                        t._file = L(k({}, t._file), {
                                            items: t.files
                                        });
                                        return [4, (0, i.V)(n)];
                                    case 1:
                                        o.sent();
                                        n.cancel();
                                        t.emit("cancel-upload-item", t._file);
                                        0 === t.files.length && t.cancel();
                                        return [2]
                                }
                            }))
                        }))()
                    };
                    r.failed = function () {
                        return this.files.some((function (e) {
                            return e.status === l.m.ERROR
                        }))
                    };
                    r._remainingUploadCount = function () {
                        var e = this.files;
                        return e.length - e.filter((function (e) {
                            return e.status === l.m.COMPLETED
                        })).length
                    };
                    r.clear = function () {
                        this.cancel();
                        this.files = []
                    };
                    return n
                }(A);

            function V(e) {
                return Y.apply(this, arguments)
            }

            function Y() {
                Y = j((function (e) {
                    var t, n, r, o = arguments;
                    return G(this, (function (i) {
                        switch (i.label) {
                            case 0:
                                t = o.length > 1 && void 0 !== o[1] && o[1], n = o.length > 2 ? o[2] : void 0;
                                r = e.map((function (e) {
                                    return new Promise((function (r, o) {
                                        switch (e.status) {
                                            case l.m.NOT_STARTED:
                                                e.upload();
                                                break;
                                            case l.m.COMPLETED:
                                                r("complete");
                                                break;
                                            case l.m.ERROR:
                                                t && e.error !== T.evJ.ENTITY_TOO_LARGE ? e.upload() : o(new Error("File failed to upload"));
                                                break;
                                            case l.m.CANCELED:
                                                o(new Error("Upload is canceled"))
                                        }
                                        e.on("complete", (function () {
                                            r("complete")
                                        }));
                                        e.on("error", (function () {
                                            o(new Error("File ".concat(e.id, " failed to upload")))
                                        }));
                                        e.on("progress", (function (e, t) {
                                            null == n || n(e, t)
                                        }))
                                    }))
                                }));
                                return [4, Promise.all(r)];
                            case 1:
                                i.sent();
                                return [2]
                        }
                    }))
                }));
                return Y.apply(this, arguments)
            }
        },
        385231: (e, t, n) => {
            n.d(t, {
                OU: () => c,
                Fi: () => p,
                li: () => f,
                KF: () => s,
                AS: () => d
            });
            n(441143);
            var r, o, i, l = n(73904),
                a = {
                    __proto__: null,
                    bg: {
                        group: " ",
                        decimal: ","
                    },
                    cs: {
                        group: " ",
                        decimal: ","
                    },
                    da: {
                        group: ".",
                        decimal: ","
                    },
                    de: {
                        group: ".",
                        decimal: ","
                    },
                    el: {
                        group: ".",
                        decimal: ","
                    },
                    "en-GB": {
                        group: ",",
                        decimal: "."
                    },
                    "en-US": {
                        group: ",",
                        decimal: "."
                    },
                    "es-ES": {
                        group: ".",
                        decimal: ","
                    },
                    fi: {
                        group: " ",
                        decimal: ","
                    },
                    fr: {
                        group: " ",
                        decimal: ","
                    },
                    hi: {
                        group: ",",
                        decimal: "."
                    },
                    hr: {
                        group: ".",
                        decimal: ","
                    },
                    hu: {
                        group: " ",
                        decimal: ","
                    },
                    it: {
                        group: ".",
                        decimal: ","
                    },
                    ja: {
                        group: ",",
                        decimal: "."
                    },
                    ko: {
                        group: ",",
                        decimal: "."
                    },
                    lt: {
                        group: " ",
                        decimal: ","
                    },
                    nl: {
                        group: ".",
                        decimal: ","
                    },
                    no: {
                        group: " ",
                        decimal: ","
                    },
                    pl: {
                        group: " ",
                        decimal: ","
                    },
                    "pt-BR": {
                        group: ".",
                        decimal: ","
                    },
                    ro: {
                        group: ".",
                        decimal: ","
                    },
                    ru: {
                        group: " ",
                        decimal: ","
                    },
                    "sv-SE": {
                        group: " ",
                        decimal: ","
                    },
                    th: {
                        group: ",",
                        decimal: "."
                    },
                    tr: {
                        group: ".",
                        decimal: ","
                    },
                    uk: {
                        group: " ",
                        decimal: ","
                    },
                    vi: {
                        group: ".",
                        decimal: ","
                    },
                    "zh-CN": {
                        group: ",",
                        decimal: "."
                    },
                    "zh-TW": {
                        group: ",",
                        decimal: "."
                    }
                },
                u = n(968696);
            n(127624);

            function c(e) {
                return null == e ? [] : e.filter((function (t, n) {
                    return "text" !== t.type || (n > 0 && n < e.length - 1 ? "" !== t.text : "" !== t.text.trim())
                }))
            }

            function s(e, t) {
                var n = e[t],
                    r = "",
                    o = !0,
                    i = !1,
                    l = void 0;
                try {
                    for (var a, u = n[Symbol.iterator](); !(o = (a = u.next()).done); o = !0) {
                        var c = a.value;
                        switch (c.type) {
                            case "text":
                            case "textMention":
                                r += c.text;
                                break;
                            case "userMention":
                                r += "<@".concat(c.userId, ">");
                                break;
                            case "channelMention":
                                r += "<#".concat(c.channelId, ">");
                                break;
                            case "roleMention":
                                r += "<@&".concat(c.roleId, ">");
                                break;
                            case "emoji":
                                r += c.surrogate;
                                break;
                            case "customEmoji":
                                r += "<".concat(c.animated ? "a" : "", ":").concat(c.name.replace(/:/g, "").split("~")[0], ":").concat(c.emojiId, ">")
                        }
                    }
                } catch (e) {
                    i = !0;
                    l = e
                } finally {
                    try {
                        o || null == u.return || u.return()
                    } finally {
                        if (i) throw l
                    }
                }
                return r
            }

            function f(e, t) {
                return null == e[t] ? null : s(e, t)
            }

            function d(e, t) {
                if (e !== i) {
                    i = e;
                    var n, l = null !== (n = a[e]) && void 0 !== n ? n : a["en-US"],
                        c = l.group,
                        s = l.decimal;
                    r = new RegExp(u.Z.escape(c), "g");
                    o = new RegExp(u.Z.escape(s), "g")
                }
                return t.replace(r, "").replace(o, ".")
            }

            function p(e, t) {
                var n, r, o = {};
                if (null == t.options) return null;
                var i = t.options;
                (null === (n = i[0]) || void 0 === n ? void 0 : n.type) === l.jw.SUB_COMMAND_GROUP && (i = i[0].options);
                (null === (r = i[0]) || void 0 === r ? void 0 : r.type) === l.jw.SUB_COMMAND && (i = i[0].options);
                var a = !0,
                    u = !1,
                    c = void 0;
                try {
                    for (var s, f = function () {
                            var t, n = s.value,
                                r = null === (t = e.options) || void 0 === t ? void 0 : t.find((function (e) {
                                    return e.name === n.name
                                }));
                            if (n.type === l.jw.ATTACHMENT) return "continue";
                            if (null == r ? void 0 : r.autocomplete) return "continue";
                            o[n.name] = n
                        }, d = i[Symbol.iterator](); !(a = (s = d.next()).done); a = !0) f()
                } catch (e) {
                    u = !0;
                    c = e
                } finally {
                    try {
                        a || null == d.return || d.return()
                    } finally {
                        if (u) throw c
                    }
                }
                return o
            }
        },
        698847: (e, t, n) => {
            n.d(t, {
                Z: () => G,
                d: () => V
            });
            var r = n(441143),
                o = n.n(r),
                i = n(744564),
                l = n(720419),
                a = n(567867),
                u = n(73904),
                c = n(769915),
                s = n(101566),
                f = n(396043),
                d = n(99885),
                p = n(245353),
                m = n(869854),
                y = n(166076),
                h = n(358624),
                v = n(671293),
                b = n(384411),
                g = n(567403),
                _ = n(255592),
                O = n(473903),
                w = n(149258),
                S = n(72580),
                E = n(124453),
                I = n(49621),
                x = n(849172),
                A = n(385231),
                T = n(977579),
                C = n(272200),
                P = n(248046),
                M = n(127624),
                j = n(2590),
                N = n(493254),
                U = n(473708);

            function R(e, t) {
                (null == t || t > e.length) && (t = e.length);
                for (var n = 0, r = new Array(t); n < t; n++) r[n] = e[n];
                return r
            }

            function k(e, t, n, r, o, i, l) {
                try {
                    var a = e[i](l),
                        u = a.value
                } catch (e) {
                    n(e);
                    return
                }
                a.done ? t(u) : Promise.resolve(u).then(r, o)
            }

            function L(e) {
                return function () {
                    var t = this,
                        n = arguments;
                    return new Promise((function (r, o) {
                        var i = e.apply(t, n);

                        function l(e) {
                            k(i, r, o, l, a, "next", e)
                        }

                        function a(e) {
                            k(i, r, o, l, a, "throw", e)
                        }
                        l(void 0)
                    }))
                }
            }

            function D(e, t) {
                return function (e) {
                    if (Array.isArray(e)) return e
                }(e) || function (e, t) {
                    var n = null == e ? null : "undefined" != typeof Symbol && e[Symbol.iterator] || e["@@iterator"];
                    if (null != n) {
                        var r, o, i = [],
                            l = !0,
                            a = !1;
                        try {
                            for (n = n.call(e); !(l = (r = n.next()).done); l = !0) {
                                i.push(r.value);
                                if (t && i.length === t) break
                            }
                        } catch (e) {
                            a = !0;
                            o = e
                        } finally {
                            try {
                                l || null == n.return || n.return()
                            } finally {
                                if (a) throw o
                            }
                        }
                        return i
                    }
                }(e, t) || function (e, t) {
                    if (!e) return;
                    if ("string" == typeof e) return R(e, t);
                    var n = Object.prototype.toString.call(e).slice(8, -1);
                    "Object" === n && e.constructor && (n = e.constructor.name);
                    if ("Map" === n || "Set" === n) return Array.from(n);
                    if ("Arguments" === n || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n)) return R(e, t)
                }(e, t) || function () {
                    throw new TypeError("Invalid attempt to destructure non-iterable instance.\\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")
                }()
            }
            var z = function (e, t) {
                    var n, r, o, i, l = {
                        label: 0,
                        sent: function () {
                            if (1 & o[0]) throw o[1];
                            return o[1]
                        },
                        trys: [],
                        ops: []
                    };
                    return i = {
                        next: a(0),
                        throw: a(1),
                        return: a(2)
                    }, "function" == typeof Symbol && (i[Symbol.iterator] = function () {
                        return this
                    }), i;

                    function a(i) {
                        return function (a) {
                            return function (i) {
                                if (n) throw new TypeError("Generator is already executing.");
                                for (; l;) try {
                                    if (n = 1, r && (o = 2 & i[0] ? r.return : i[0] ? r.throw || ((o = r.return) && o.call(r), 0) : r.next) && !(o = o.call(r, i[1])).done) return o;
                                    (r = 0, o) && (i = [2 & i[0], o.value]);
                                    switch (i[0]) {
                                        case 0:
                                        case 1:
                                            o = i;
                                            break;
                                        case 4:
                                            l.label++;
                                            return {
                                                value: i[1], done: !1
                                            };
                                        case 5:
                                            l.label++;
                                            r = i[1];
                                            i = [0];
                                            continue;
                                        case 7:
                                            i = l.ops.pop();
                                            l.trys.pop();
                                            continue;
                                        default:
                                            if (!(o = l.trys, o = o.length > 0 && o[o.length - 1]) && (6 === i[0] || 2 === i[0])) {
                                                l = 0;
                                                continue
                                            }
                                            if (3 === i[0] && (!o || i[1] > o[0] && i[1] < o[3])) {
                                                l.label = i[1];
                                                break
                                            }
                                            if (6 === i[0] && l.label < o[1]) {
                                                l.label = o[1];
                                                o = i;
                                                break
                                            }
                                            if (o && l.label < o[2]) {
                                                l.label = o[2];
                                                l.ops.push(i);
                                                break
                                            }
                                            o[2] && l.ops.pop();
                                            l.trys.pop();
                                            continue
                                    }
                                    i = t.call(e, l)
                                } catch (e) {
                                    i = [6, e];
                                    r = 0
                                } finally {
                                    n = o = 0
                                }
                                if (5 & i[0]) throw i[1];
                                return {
                                    value: i[0] ? i[1] : void 0,
                                    done: !0
                                }
                            }([i, a])
                        }
                    }
                },
                Z = function (e, t) {
                    var n;
                    return null === (n = null == e ? void 0 : e.find((function (e) {
                        return e.displayName === t
                    }))) || void 0 === n ? void 0 : n.value
                },
                B = function (e, t, n) {
                    var r, o = arguments.length > 3 && void 0 !== arguments[3] ? arguments[3] : function (e) {
                            return e
                        },
                        i = e.name === (null === (r = n.autocomplete) || void 0 === r ? void 0 : r.name);
                    if (i) return n.autocomplete.query;
                    if ("" === t) return null;
                    var l = x.Z.getAutocompleteLastChoices(n.channel.id, e.name);
                    if (null != l) {
                        var a;
                        return null !== (a = Z(l, t)) && void 0 !== a ? a : o(t)
                    }
                    return o(t)
                },
                F = function (e) {
                    var t = e.toLowerCase() === M.WO.toLowerCase(),
                        n = e.toLowerCase() === M.Lu.toLowerCase();
                    return t || n ? t : null
                };

            function G(e) {
                return H.apply(this, arguments)
            }

            function H() {
                return (H = L((function (e) {
                    var t, n, r, l, c, s, p, m, y, h, g, O, w, S, E, x, T, M, N, U, R, k, L, D, G, H, V, K, $, q, W, X, Q, ee, te, ne, re, oe, ie, le, ae, ue, ce, se, fe, de, pe, me, ye, he;
                    return z(this, (function (z) {
                        switch (z.label) {
                            case 0:
                                t = e.command, n = e.optionValues, r = e.context, l = e.commandTargetId, c = e.maxSizeCallback, s = e.commandOrigin, p = void 0 === s ? C.bB.CHAT : s;
                                null == r.autocomplete && i.Z.dispatch({
                                    type: "APPLICATION_COMMAND_USED",
                                    context: r,
                                    command: t
                                });
                                return [4, v.Z.unarchiveThreadIfNecessary(r.channel.id)];
                            case 1:
                                z.sent();
                                h = [];
                                g = [];
                                O = (0, P.D7)(p);
                                if (null != t.options) {
                                    w = !0, S = !1, E = void 0;
                                    try {
                                        for (x = t.options[Symbol.iterator](); !(w = (T = x.next()).done); w = !0)
                                            if ((M = T.value).type !== u.jw.SUB_COMMAND && M.type !== u.jw.SUB_COMMAND_GROUP && M.name in n) {
                                                U = (null === (N = r.autocomplete) || void 0 === N ? void 0 : N.name) === M.name || void 0;
                                                R = null;
                                                if (M.type !== u.jw.STRING)
                                                    if (M.type !== u.jw.ATTACHMENT) {
                                                        V = A.OU(n[M.name]);
                                                        o()(null != r.autocomplete || 1 === V.length, 'Option "'.concat(M.name, '" expects a single option type'));
                                                        if (null != V[0] || U) {
                                                            $ = null !== (K = V[0]) && void 0 !== K ? K : {
                                                                type: "text",
                                                                text: ""
                                                            };
                                                            switch (M.type) {
                                                                case u.jw.CHANNEL:
                                                                    if ("channelMention" === $.type) R = $.channelId;
                                                                    else if ("text" === $.type)
                                                                        if ((0, P.BH)($.text)) R = $.text.trim();
                                                                        else {
                                                                            W = (0, d.K)($.text, null === (q = r.guild) || void 0 === q ? void 0 : q.id, r.channel.id);
                                                                            o()("channelMention" === (null == W ? void 0 : W.type), "Failed to resolve ".concat($.text));
                                                                            R = W.channelId
                                                                        } break;
                                                                case u.jw.ROLE:
                                                                    if ("roleMention" === $.type) R = $.roleId;
                                                                    else if ("text" === $.type)
                                                                        if ((0, P.BH)($.text)) R = $.text.trim();
                                                                        else {
                                                                            Q = (0,
                                                                                d.K)($.text, null === (X = r.guild) || void 0 === X ? void 0 : X.id, r.channel.id, {
                                                                                allowUsers: !1
                                                                            });
                                                                            o()("roleMention" === (null == Q ? void 0 : Q.type), "Failed to resolve ".concat($.text));
                                                                            R = Q.roleId
                                                                        }
                                                                    else "textMention" === $.type && "@everyone" === $.text && (R = null === (ee = r.guild) || void 0 === ee ? void 0 : ee.id);
                                                                    break;
                                                                case u.jw.USER:
                                                                    if ("userMention" === $.type) R = $.userId;
                                                                    else if ("text" === $.type)
                                                                        if ((0, P.BH)($.text)) R = $.text.trim();
                                                                        else {
                                                                            ne = (0, d.K)($.text, null === (te = r.guild) || void 0 === te ? void 0 : te.id, r.channel.id, {
                                                                                allowRoles: !1
                                                                            });
                                                                            o()("userMention" === (null == ne ? void 0 : ne.type), "Failed to resolve ".concat($.text));
                                                                            R = ne.userId
                                                                        } break;
                                                                case u.jw.MENTIONABLE:
                                                                    "userMention" === $.type ? R = $.userId : "roleMention" === $.type ? R = $.roleId : "textMention" === $.type && "@everyone" === $.text ? R = null === (re = r.guild) || void 0 === re ? void 0 : re.id : "text" === $.type && ((0, P.BH)($.text) ? R = $.text.trim() : "userMention" === (null == (ie = (0,
                                                                        d.K)($.text, null === (oe = r.guild) || void 0 === oe ? void 0 : oe.id, r.channel.id)) ? void 0 : ie.type) ? R = ie.userId : "roleMention" === (null == ie ? void 0 : ie.type) ? R = ie.roleId : "textMention" === (null == ie ? void 0 : ie.type) && "@everyone" === ie.text ? R = null === (le = r.guild) || void 0 === le ? void 0 : le.id : o()(!1, "Failed to resolve ".concat($.text)));
                                                                    break;
                                                                case u.jw.BOOLEAN:
                                                                    "text" === $.type && (R = F($.text.trim()));
                                                                    break;
                                                                case u.jw.INTEGER:
                                                                    if ("text" === $.type) {
                                                                        ae = $.text.trim();
                                                                        R = null != M.choices ? Number(Z(M.choices, ae)) : M.autocomplete ? B(M, ae, r, Number) : Number(A.AS(b.default.locale, ae))
                                                                    }
                                                                    break;
                                                                case u.jw.NUMBER:
                                                                    if ("text" === $.type) {
                                                                        ue = $.text.trim();
                                                                        R = null != M.choices ? Number(Z(M.choices, ue)) : M.autocomplete ? B(M, ue, r, Number) : Number(A.AS(b.default.locale, ue))
                                                                    }
                                                                    break;
                                                                default:
                                                                    o()(!1, "Unsupported option type: ".concat(M.type));
                                                                    continue
                                                            }
                                                            o()(null != r.autocomplete || null != R, 'Unexpected value for option "'.concat(M.name, '"'));
                                                            null != R && h.push({
                                                                type: M.type,
                                                                name: M.name,
                                                                value: R,
                                                                focused: U
                                                            })
                                                        }
                                                    } else {
                                                        if (null != r.autocomplete) continue;
                                                        if (null == (G = _.Z.getUpload(r.channel.id, M.name, O))) continue;
                                                        H = _.Z.getUploads(r.channel.id, O).findIndex((function (e) {
                                                            return G.id === e.id
                                                        }));
                                                        g.push(G);
                                                        R = H;
                                                        h.push({
                                                            type: M.type,
                                                            name: M.name,
                                                            value: R,
                                                            focused: U
                                                        })
                                                    }
                                                else {
                                                    D = null !== (L = null === (k = A.li(n, M.name)) || void 0 === k ? void 0 : k.trim()) && void 0 !== L ? L : "";
                                                    R = null != M.choices ? Z(M.choices, D) : M.autocomplete ? B(M, D, r) : D;
                                                    o()(null != r.autocomplete || null != R, 'Option "'.concat(M.name, '" expects a value'));
                                                    null != R && h.push({
                                                        type: M.type,
                                                        name: M.name,
                                                        value: R,
                                                        focused: U
                                                    })
                                                }
                                            }
                                    } catch (e) {
                                        S = !0;
                                        E = e
                                    } finally {
                                        try {
                                            w || null == x.return || x.return()
                                        } finally {
                                            if (S) throw E
                                        }
                                    }
                                }
                                if (null != t.subCommandPath)
                                    for (ce = t.subCommandPath.length - 1; ce >= 0; ce -= 1) {
                                        se = t.subCommandPath[ce], fe = se.name, de = se.type;
                                        h = [{
                                            type: de,
                                            name: fe,
                                            options: h
                                        }]
                                    }
                                if (null != t.execute) {
                                    f.ZP.trackWithMetadata(j.rMx.APPLICATION_COMMAND_USED, {
                                        command_id: t.id,
                                        application_id: t.applicationId,
                                        command_type: t.type
                                    });
                                    return [2, t.execute(h, r)]
                                }
                                if (t.inputType === C.iw.BUILT_IN || t.inputType === C.iw.BUILT_IN_TEXT || t.inputType === C.iw.BUILT_IN_INTEGRATION) return [2];
                                ye = {
                                    version: t.version,
                                    id: null !== (pe = null === (m = t.rootCommand) || void 0 === m ? void 0 : m.id) && void 0 !== pe ? pe : t.id,
                                    guild_id: t.guildId,
                                    name: null !== (me = null === (y = t.rootCommand) || void 0 === y ? void 0 : y.name) && void 0 !== me ? me : t.name,
                                    type: t.type,
                                    options: h,
                                    application_command: t.rootCommand
                                };
                                he = function () {
                                    J(n)
                                };
                                null != l && (ye.target_id = l);
                                if (null != r.autocomplete)(0, I.GV)(t, r, ye);
                                else {
                                    a.Z.clearAll(r.channel.id, O);
                                    Y({
                                        applicationId: t.applicationId,
                                        data: ye,
                                        context: r,
                                        attachments: g,
                                        maxSizeCallback: c,
                                        onMessageSuccess: he,
                                        commandDisplayName: t.displayName
                                    })
                                }
                                return [2]
                        }
                    }))
                }))).apply(this, arguments)
            }
            var J = function (e) {
                    var t = Object.values(e).flatMap((function (e) {
                        return e.map((function (e) {
                            return "emoji" === e.type ? {
                                name: e.name.replaceAll(":", "")
                            } : "customEmoji" === e.type ? p.Z.getCustomEmojiById(e.emojiId) : null
                        })).filter(S.lm)
                    }));
                    t.length > 0 && i.Z.dispatch({
                        type: "EMOJI_TRACK_USAGE",
                        emojiUsed: t
                    })
                },
                V = function (e, t, n) {
                    if (e.isCommandType()) {
                        var r = t.guild_id;
                        null != e.interactionData && Y({
                            applicationId: n,
                            data: e.interactionData,
                            context: {
                                channel: t,
                                guild: null != r ? g.Z.getGuild(r) : null
                            }
                        })
                    }
                },
                Y = function (e) {
                    var t = e.applicationId,
                        n = e.data,
                        r = e.context,
                        o = e.attachments,
                        a = e.maxSizeCallback,
                        c = e.onMessageSuccess,
                        s = e.commandDisplayName,
                        f = r.channel,
                        d = r.guild,
                        p = f.id,
                        y = null == d ? void 0 : d.id,
                        v = D(function (e, t, n) {
                            var r, o = null === (r = T.ZP.getApplicationSections(e.channel.id, n)) || void 0 === r ? void 0 : r.find((function (e) {
                                return e.id === t
                            }));
                            if (null != o) {
                                var l, a, u, c = null !== (a = null === (l = o.application) || void 0 === l ? void 0 : l.bot) && void 0 !== a ? a : {
                                    id: o.id,
                                    username: o.name,
                                    discriminator: "0000",
                                    avatar: null,
                                    bot: !0
                                };
                                i.Z.dispatch({
                                    type: "STORE_APPLICATION_INTERACTION_FAKE_USER",
                                    user: c
                                });
                                return [c, null !== (u = o.application) && void 0 !== u ? u : null]
                            }
                            return [null, null]
                        }(r, t, n.type), 2),
                        b = v[0],
                        g = v[1],
                        _ = n.type === u.yU.CHAT ? j.uaV.CHAT_INPUT_COMMAND : j.uaV.CONTEXT_MENU_COMMAND,
                        w = (0, h.ZP)({
                            channelId: p,
                            content: "",
                            tts: !1,
                            type: _,
                            messageReference: void 0,
                            allowedMentions: void 0,
                            author: null != b ? b : void 0
                        });
                    w.application = null != g ? g : void 0;
                    w.interaction = {
                        id: n.id,
                        name: n.name,
                        name_localized: s,
                        type: u.B8.APPLICATION_COMMAND,
                        user: (0, h.pe)(O.default.getCurrentUser())
                    };
                    w.interaction_data = n;
                    var S = {
                        applicationId: t,
                        channelId: p,
                        guildId: y,
                        data: n,
                        nonce: w.id,
                        attachments: o,
                        maxSizeCallback: a
                    };
                    l.Z.receiveMessage(p, w, !0, {
                        applicationId: t
                    });
                    m.kz(S.nonce, {
                        messageId: w.id,
                        onCreate: function (e) {
                            null != w.interaction && (w.interaction.id = e)
                        },
                        onFailure: function (e, t) {
                            return function (e, t) {
                                null == t && null != e && l.Z.sendClydeError(p, e);
                                i.Z.dispatch({
                                    type: "MESSAGE_SEND_FAILED",
                                    messageId: w.id,
                                    channelId: p,
                                    reason: t
                                })
                            }(e, t)
                        },
                        data: {
                            interactionType: u.B8.APPLICATION_COMMAND,
                            channelId: p
                        }
                    });
                    if (null != o) {
                        (function (e, t, n, r) {
                            return W.apply(this, arguments)
                        })(o, S.nonce, y, a).then((function (e) {
                            e && K(S, c)
                        }))
                    } else K(S, c)
                };

            function K(e, t) {
                c.ZP.enqueue({
                    type: c.$V.COMMAND,
                    message: e
                }, (function (n) {
                    (0, y.Sg)(e.nonce, n);
                    n.ok && null != t && t()
                }))
            }

            function $(e, t) {
                return q.apply(this, arguments)
            }

            function q() {
                return (q = L((function (e, t) {
                    var n, r, o, i, l, a, u, c, s, f, d, p;
                    return z(this, (function (m) {
                        switch (m.label) {
                            case 0:
                                n = 0;
                                r = 0;
                                o = !0, i = !1, l = void 0;
                                m.label = 1;
                            case 1:
                                m.trys.push([1, 8, 9, 10]);
                                a = e[Symbol.iterator]();
                                m.label = 2;
                            case 2:
                                if (o = (u = a.next()).done) return [3, 7];
                                c = u.value;
                                if (!t) return [3, 3];
                                d = null !== (s = c.currentSize) && void 0 !== s ? s : 0;
                                return [3, 5];
                            case 3:
                                return [4, c.getSize()];
                            case 4:
                                d = m.sent();
                                m.label = 5;
                            case 5:
                                (f = d) > r && (r = f);
                                n += f;
                                m.label = 6;
                            case 6:
                                o = !0;
                                return [3, 2];
                            case 7:
                                return [3, 10];
                            case 8:
                                p = m.sent();
                                i = !0;
                                l = p;
                                return [3, 10];
                            case 9:
                                try {
                                    o || null == a.return || a.return()
                                } finally {
                                    if (i) throw l
                                }
                                return [7];
                            case 10:
                                return [2, {
                                    totalSize: n,
                                    largestUploadedFileSize: r
                                }]
                        }
                    }))
                }))).apply(this, arguments)
            }

            function W() {
                return (W = L((function (e, t, n, r) {
                    var o, i, l, a, u, c, f;
                    return z(this, (function (d) {
                        switch (d.label) {
                            case 0:
                                o = e;
                                i = (0, w.dg)(n);
                                l = function (e) {
                                    null == r || r(i, e);
                                    m.yr(t, j.evJ.ENTITY_TOO_LARGE, U.Z.Messages.UPLOAD_AREA_TOO_LARGE_HELP.format({
                                        maxSize: (0, w.Ng)(i)
                                    }))
                                };
                                return [4, $(o, !1)];
                            case 1:
                                a = d.sent(), u = a.totalSize;
                                if ((c = a.largestUploadedFileSize) > Math.max(i, N.Y1) || u > E.zz) {
                                    l(c);
                                    return [2, !1]
                                }
                                d.label = 2;
                            case 2:
                                d.trys.push([2, 4, , 5]);
                                return [4, (0, s.$)(o)];
                            case 3:
                                d.sent();
                                return [3, 5];
                            case 4:
                                d.sent();
                                m.yr(t, void 0, U.Z.Messages.UPLOADING_FILES_FAILED.format({
                                    count: o.length
                                }));
                                return [3, 5];
                            case 5:
                                return [4, $(o, !0)];
                            case 6:
                                f = d.sent(), u = f.totalSize, c = f.largestUploadedFileSize;
                                if (o.some((function (e) {
                                        return e.error === j.evJ.ENTITY_TOO_LARGE
                                    })) || u > E.zz) {
                                    l(c);
                                    return [2, !1]
                                }
                                return [2, !0]
                        }
                    }))
                }))).apply(this, arguments)
            }
        },
        99885: (e, t, n) => {
            n.d(t, {
                i: () => b,
                K: () => g
            });
            var r = n(773011),
                o = n(23925),
                i = n(245353),
                l = n(135855),
                a = n(18882),
                u = n(61209),
                c = n(5544),
                s = n(21372),
                f = n(567403),
                d = n(473903),
                p = n(504275),
                m = n(855483),
                y = n(876986);

            function h(e, t) {
                (null == t || t > e.length) && (t = e.length);
                for (var n = 0, r = new Array(t); n < t; n++) r[n] = e[n];
                return r
            }

            function v(e, t) {
                return function (e) {
                    if (Array.isArray(e)) return e
                }(e) || function (e, t) {
                    var n = null == e ? null : "undefined" != typeof Symbol && e[Symbol.iterator] || e["@@iterator"];
                    if (null != n) {
                        var r, o, i = [],
                            l = !0,
                            a = !1;
                        try {
                            for (n = n.call(e); !(l = (r = n.next()).done); l = !0) {
                                i.push(r.value);
                                if (t && i.length === t) break
                            }
                        } catch (e) {
                            a = !0;
                            o = e
                        } finally {
                            try {
                                l || null == n.return || n.return()
                            } finally {
                                if (a) throw o
                            }
                        }
                        return i
                    }
                }(e, t) || function (e, t) {
                    if (!e) return;
                    if ("string" == typeof e) return h(e, t);
                    var n = Object.prototype.toString.call(e).slice(8, -1);
                    "Object" === n && e.constructor && (n = e.constructor.name);
                    if ("Map" === n || "Set" === n) return Array.from(n);
                    if ("Arguments" === n || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n)) return h(e, t)
                }(e, t) || function () {
                    throw new TypeError("Invalid attempt to destructure non-iterable instance.\\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")
                }()
            }

            function b(e, t, n, d) {
                var m = null != d ? d : {},
                    h = m.allowUsers,
                    b = void 0 === h || h,
                    g = m.allowRoles,
                    O = void 0 === g || g;
                switch (e[0]) {
                    case "@":
                        return function (e, t, n, r, i) {
                            var l = v(e.slice(1).split("#", 2), 2),
                                a = l[0],
                                c = l[1],
                                d = null != t ? f.Z.getGuild(t) : null,
                                p = (0, o.M9)(d);
                            if (i && null == c && null != d) {
                                var m = !0,
                                    h = !1,
                                    b = void 0;
                                try {
                                    for (var g, O = Object.values(d.roles)[Symbol.iterator](); !(m = (g = O.next()).done); m = !0) {
                                        var w = g.value;
                                        if (a === w.name) return {
                                            type: "roleMention",
                                            roleId: w.id,
                                            children: [{
                                                text: ""
                                            }]
                                        }
                                    }
                                } catch (e) {
                                    h = !0;
                                    b = e
                                } finally {
                                    try {
                                        m || null == O.return || O.return()
                                    } finally {
                                        if (h) throw b
                                    }
                                }
                            }
                            if (r) {
                                var S = null != n ? u.Z.getChannel(n) : null;
                                if (null != S)
                                    if (S.isPrivate()) {
                                        var E = !0,
                                            I = !1,
                                            x = void 0;
                                        try {
                                            for (var A, T = S.recipients[Symbol.iterator](); !(E = (A = T.next()).done); E = !0) {
                                                var C = A.value;
                                                if (_(a, c, C)) return {
                                                    type: "userMention",
                                                    userId: C,
                                                    children: [{
                                                        text: ""
                                                    }]
                                                }
                                            }
                                        } catch (e) {
                                            I = !0;
                                            x = e
                                        } finally {
                                            try {
                                                E || null == T.return || T.return()
                                            } finally {
                                                if (I) throw x
                                            }
                                        }
                                    } else {
                                        var P = s.ZP.getMembers(t),
                                            M = !0,
                                            j = !1,
                                            N = void 0;
                                        try {
                                            for (var U, R = P[Symbol.iterator](); !(M = (U = R.next()).done); M = !0) {
                                                var k = U.value.userId;
                                                if (_(a, c, k)) return {
                                                    type: "userMention",
                                                    userId: k,
                                                    children: [{
                                                        text: ""
                                                    }]
                                                }
                                            }
                                        } catch (e) {
                                            j = !0;
                                            N = e
                                        } finally {
                                            try {
                                                M || null == R.return || R.return()
                                            } finally {
                                                if (j) throw N
                                            }
                                        }
                                        if (p && _(a, c, y.fL)) return {
                                            type: "userMention",
                                            userId: y.fL,
                                            children: [{
                                                text: ""
                                            }]
                                        }
                                    }
                            }
                            return null
                        }(e, t, n, b, O);
                    case ":":
                        return function (e, t) {
                            var n = l.ZP.EMOJI_NAME_RE.exec(e);
                            if (null == n) return null;
                            var r = n[1],
                                o = i.Z.getDisambiguatedEmojiContext(t).getCustomEmoji();
                            if (null != o && r in o) {
                                var a = o[r];
                                return {
                                    type: "customEmoji",
                                    emoji: {
                                        emojiId: a.id,
                                        name: "require_colons" in a && a.require_colons ? ":".concat(a.name, ":") : a.name,
                                        animated: !0 === a.animated,
                                        jumboable: !1
                                    },
                                    children: [{
                                        text: ""
                                    }]
                                }
                            }
                            return null
                        }(e, t);
                    case "#":
                        return function (e, t) {
                            if (null == t) return null;
                            var n;
                            n = e.length > 3 && '"' === e[1] && '"' === e[e.length - 1] ? (0, r.mA)(e.slice(2, e.length - 1)) : e.slice(1);
                            var o = c.ZP.getTextChannelNameDisambiguations(t),
                                i = !0,
                                l = !1,
                                u = void 0;
                            try {
                                for (var s, f = Object.keys(o)[Symbol.iterator](); !(i = (s = f.next()).done); i = !0) {
                                    var d = s.value;
                                    if (o[d].name === n) return {
                                        type: "channelMention",
                                        channelId: d,
                                        children: [{
                                            text: ""
                                        }]
                                    }
                                }
                            } catch (e) {
                                l = !0;
                                u = e
                            } finally {
                                try {
                                    i || null == f.return || f.return()
                                } finally {
                                    if (l) throw u
                                }
                            }
                            var m = !0,
                                y = !1,
                                h = void 0;
                            try {
                                for (var v, b = p.k1[Symbol.iterator](); !(m = (v = b.next()).done); m = !0) {
                                    var g = v.value;
                                    if (g !== c.sH) {
                                        var _ = c.ZP.getChannels(t)[g],
                                            O = !0,
                                            w = !1,
                                            S = void 0;
                                        try {
                                            for (var E, I = _[Symbol.iterator](); !(O = (E = I.next()).done); O = !0) {
                                                var x = E.value.channel;
                                                if (x.name === n) return {
                                                    type: "channelMention",
                                                    channelId: x.id,
                                                    children: [{
                                                        text: ""
                                                    }]
                                                }
                                            }
                                        } catch (e) {
                                            w = !0;
                                            S = e
                                        } finally {
                                            try {
                                                O || null == I.return || I.return()
                                            } finally {
                                                if (w) throw S
                                            }
                                        }
                                    }
                                }
                            } catch (e) {
                                y = !0;
                                h = e
                            } finally {
                                try {
                                    m || null == b.return || b.return()
                                } finally {
                                    if (y) throw h
                                }
                            }
                            var A = a.Z.getActiveJoinedThreadsForGuild(t),
                                T = !0,
                                C = !1,
                                P = void 0;
                            try {
                                for (var M, j = Object.keys(A)[Symbol.iterator](); !(T = (M = j.next()).done); T = !0) {
                                    var N = M.value,
                                        U = !0,
                                        R = !1,
                                        k = void 0;
                                    try {
                                        for (var L, D = Object.keys(A[N])[Symbol.iterator](); !(U = (L = D.next()).done); U = !0) {
                                            var z = L.value,
                                                Z = A[N][z].channel;
                                            if (Z.name === n) return {
                                                type: "channelMention",
                                                channelId: Z.id,
                                                children: [{
                                                    text: ""
                                                }]
                                            }
                                        }
                                    } catch (e) {
                                        R = !0;
                                        k = e
                                    } finally {
                                        try {
                                            U || null == D.return || D.return()
                                        } finally {
                                            if (R) throw k
                                        }
                                    }
                                }
                            } catch (e) {
                                C = !0;
                                P = e
                            } finally {
                                try {
                                    T || null == j.return || j.return()
                                } finally {
                                    if (C) throw P
                                }
                            }
                            return null
                        }(e, t)
                }
                return null
            }

            function g(e, t, n, r) {
                var o = b(e, t, n, r);
                return null == o ? null : (0, m.VI)(o)
            }

            function _(e, t, n) {
                var r = d.default.getUser(n);
                return null != r && (n === y.fL && "clyde" === e.toLowerCase() || r.username === e && r.discriminator === (null != t ? t : "0"))
            }
        }
    }
]);
//# sourceMappingURL=03a0e74444c874859a9e.js.map